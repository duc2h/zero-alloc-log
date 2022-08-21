# zero-alloc-log

This is my demo about zero allocation in logging.

#### Poollog result
##### Benchmark
```
go test -cpu=1,2,5 -benchmem -benchtime=5s -bench . ./poollog/...
goos: linux
goarch: amd64
pkg: zero-log/poollog
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkEvent                          217363773               33.56 ns/op            0 B/op          0 allocs/op
BenchmarkEvent-2                        162484694               39.37 ns/op            0 B/op          0 allocs/op
BenchmarkEvent-5                        159550320               36.57 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithDone                  174325942               34.17 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithDone-2                161502086               38.18 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithDone-5                179667643               33.44 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithInterface             197521564               31.37 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithInterface-2           166555442               36.21 ns/op            0 B/op          0 allocs/op
BenchmarkEventWithInterface-5           168423651               35.02 ns/op            0 B/op          0 allocs/op
PASS
ok      zero-log/poollog        87.047s
```
##### Escape to heap
```
go build -gcflags '-m' ./poollog/...
# zero-log/poollog
poollog/event.go:37:7: can inline glob..func1
poollog/event.go:53:6: can inline putEvent
poollog/event.go:61:10: inlining call to putEvent
poollog/event.go:68:10: inlining call to putEvent
poollog/event.go:75:10: inlining call to putEvent
poollog/event.go:82:10: inlining call to putEvent
poollog/event.go:92:6: can inline fatalFunc
poollog/event.go:89:10: inlining call to putEvent
poollog/event.go:101:6: can inline panicFunc
poollog/event.go:98:10: inlining call to putEvent
poollog/event.go:107:10: inlining call to putEvent
poollog/interface.go:7:6: can inline newEventInterface
poollog/interface.go:11:6: can inline putEventInterface
poollog/interface.go:20:24: inlining call to newEventInterface
poollog/interface.go:22:19: inlining call to putEventInterface
poollog/interface.go:21:9: devirtualizing e.write to *event
poollog/event.go:28:7: leaking param: e
poollog/event.go:38:10: &event{...} escapes to heap
poollog/event.go:39:13: make([]byte, 0, 500) escapes to heap
poollog/event.go:46:8: newEvent ignoring self-assignment in e.buf = e.buf[:0]
poollog/event.go:44:15: buf does not escape
poollog/event.go:44:40: leaking param: done
poollog/event.go:53:15: leaking param: e
poollog/event.go:58:12: msg does not escape
poollog/event.go:59:22: ([]byte)(msg) does not escape
poollog/event.go:65:11: msg does not escape
poollog/event.go:66:22: ([]byte)(msg) does not escape
poollog/event.go:72:11: msg does not escape
poollog/event.go:73:22: ([]byte)(msg) does not escape
poollog/event.go:79:12: leaking param: err
poollog/event.go:80:22: ([]byte)(err.Error()) does not escape
poollog/event.go:92:16: b does not escape
poollog/event.go:86:12: leaking param: err
poollog/event.go:87:22: ([]byte)(err.Error()) does not escape
poollog/event.go:101:16: leaking param: b
poollog/event.go:101:34: b escapes to heap
poollog/event.go:95:12: leaking param: err
poollog/event.go:96:22: ([]byte)(err.Error()) does not escape
poollog/event.go:104:19: msg does not escape
poollog/event.go:104:31: leaking param: done
poollog/event.go:105:22: ([]byte)(msg) does not escape
poollog/interface.go:7:24: buf does not escape
poollog/interface.go:7:49: leaking param: done
poollog/interface.go:11:24: leaking param: e
poollog/interface.go:13:8: "invalid type" escapes to heap
poollog/interface.go:19:22: msg does not escape
poollog/interface.go:20:31: ([]byte)(msg) does not escape
poollog/interface.go:22:19: "invalid type" escapes to heap
```
#### Normallog result
##### Benchmark
```
go test -cpu=1,2,5 -benchmem -benchtime=5s -bench . ./normallog/...
goos: linux
goarch: amd64
pkg: zero-log/normallog
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkEvent                          291643672               31.78 ns/op            8 B/op          1 allocs/op
BenchmarkEvent-2                        176535709               36.28 ns/op            8 B/op          1 allocs/op
BenchmarkEvent-5                        198817448               30.39 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithDone                  186625639               31.02 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithDone-2                162883490               36.42 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithDone-5                178659344               32.04 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithInterface             249911389               22.58 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithInterface-2           186407341               34.10 ns/op            8 B/op          1 allocs/op
BenchmarkEventWithInterface-5           187671218               30.92 ns/op            8 B/op          1 allocs/op
PASS
ok      zero-log/normallog      84.845s
```
##### Escape to heap
```
# zero-log/normallog
normallog/event.go:35:6: can inline newEvent
normallog/event.go:43:6: can inline Debug
normallog/event.go:44:15: inlining call to newEvent
normallog/event.go:48:6: can inline Info
normallog/event.go:49:15: inlining call to newEvent
normallog/event.go:53:6: can inline Warn
normallog/event.go:54:15: inlining call to newEvent
normallog/event.go:59:15: inlining call to newEvent
normallog/event.go:64:6: can inline fatalFunc
normallog/event.go:69:15: inlining call to newEvent
normallog/event.go:74:6: can inline panicFunc
normallog/event.go:79:15: inlining call to newEvent
normallog/event.go:84:6: can inline InfoWithDone
normallog/event.go:85:15: inlining call to newEvent
normallog/interface.go:7:6: can inline newEventInterface
normallog/interface.go:8:17: inlining call to newEvent
normallog/interface.go:11:6: can inline InforByInterface
normallog/interface.go:12:15: inlining call to newEvent
normallog/event.go:27:7: leaking param: e
normallog/event.go:35:15: leaking param: buf
normallog/event.go:35:40: leaking param: done
normallog/event.go:36:9: &event{...} escapes to heap
normallog/event.go:43:12: msg does not escape
normallog/event.go:44:22: ([]byte)(msg) escapes to heap
normallog/event.go:44:15: &event{...} does not escape
normallog/event.go:48:11: msg does not escape
normallog/event.go:49:22: ([]byte)(msg) escapes to heap
normallog/event.go:49:15: &event{...} does not escape
normallog/event.go:53:11: msg does not escape
normallog/event.go:54:22: ([]byte)(msg) escapes to heap
normallog/event.go:54:15: &event{...} does not escape
normallog/event.go:58:12: leaking param: err
normallog/event.go:59:22: ([]byte)(err.Error()) escapes to heap
normallog/event.go:59:15: &event{...} does not escape
normallog/event.go:64:16: b does not escape
normallog/event.go:68:12: leaking param: err
normallog/event.go:69:22: ([]byte)(err.Error()) escapes to heap
normallog/event.go:69:15: &event{...} does not escape
normallog/event.go:74:16: leaking param: b
normallog/event.go:75:8: b escapes to heap
normallog/event.go:78:12: leaking param: err
normallog/event.go:79:22: ([]byte)(err.Error()) escapes to heap
normallog/event.go:79:15: &event{...} does not escape
normallog/event.go:84:19: msg does not escape
normallog/event.go:84:31: leaking param: done
normallog/event.go:85:22: ([]byte)(msg) escapes to heap
normallog/event.go:85:15: &event{...} does not escape
normallog/interface.go:7:24: leaking param: buf
normallog/interface.go:7:49: leaking param: done
normallog/interface.go:8:17: &event{...} escapes to heap
normallog/interface.go:11:23: msg does not escape
normallog/interface.go:12:22: ([]byte)(msg) escapes to heap
normallog/interface.go:12:15: &event{...} does not escape
```