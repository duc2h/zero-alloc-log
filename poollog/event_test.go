package poollog

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStream struct {
	writeCount uint64
}

func (t *testStream) WriteCount() uint64 {
	return atomic.LoadUint64(&t.writeCount)
}

func (t *testStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&t.writeCount, 1)
	return len(p), nil
}

func done(buf []byte) {}

func TestNewEvent(t *testing.T) {
	stream := &testStream{}
	LogWriter = stream
	Info("test")

	t.Log(stream.writeCount)
	assert.Equal(t, 1, int(stream.WriteCount()))
}

func BenchmarkEvent(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("test")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatal("Log write count")
	}
}

func BenchmarkEventWithDone(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			InfoWithDone("test", done)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatal("Log write count")
	}
}
