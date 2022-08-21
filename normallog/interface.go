package normallog

type eventInterface interface {
	write()
}

func newEventInterface(buf []byte, level Level, done func(b []byte)) eventInterface {
	return newEvent(buf, level, done)
}

func InforByInterface(msg string) {
	e := newEvent([]byte(msg), InfoLevel, nil)
	e.write()
}
