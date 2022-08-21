package normallog

import (
	"io"
	"os"
)

var LogWriter io.Writer

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

type event struct {
	buf   []byte         // log message
	level Level          // log level
	done  func(b []byte) // after writing (for Panic, Fatal)
}

func (e event) write() {
	if e.done != nil {
		defer e.done(e.buf)
	}

	LogWriter.Write(e.buf)
}

func newEvent(buf []byte, level Level, done func(b []byte)) *event {
	return &event{
		buf:   buf,
		level: level,
		done:  done,
	}
}

func Debug(msg string) {
	e := newEvent([]byte(msg), DebugLevel, nil)
	e.write()
}

func Info(msg string) {
	e := newEvent([]byte(msg), InfoLevel, nil)
	e.write()
}

func Warn(msg string) {
	e := newEvent([]byte(msg), WarnLevel, nil)
	e.write()
}

func Error(err error) {
	e := newEvent([]byte(err.Error()), ErrorLevel, nil)
	e.write()
}

// fatal log
func fatalFunc(b []byte) {
	os.Exit(1)
}

func Fatal(err error) {
	e := newEvent([]byte(err.Error()), FatalLevel, fatalFunc)
	e.write()
}

// panic log
func panicFunc(b []byte) {
	panic(b)
}

func Panic(err error) {
	e := newEvent([]byte(err.Error()), PanicLevel, panicFunc)
	e.write()
}

// InfoWithDone is the func for bench test of using channel.
func InfoWithDone(msg string, done func(b []byte)) {
	e := newEvent([]byte(msg), InfoLevel, done)
	e.write()
}
