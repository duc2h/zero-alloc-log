package normallog

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStream struct {
	writeCount uint64
}

func (s *testStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *testStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}

func doneFunc(b []byte) {}

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
		b.Fatalf("Log write count")
	}
}

func BenchmarkEventWithDone(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			InfoWithDone("test", doneFunc)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
