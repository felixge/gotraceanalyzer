# gotraceanlyzer

Command gotraceanalyzer turns golang tracebacks into useful summaries.

If you ever starred at a go panic that showed thousands of active goroutines,
this tool will help you to make more sense of it.

Summaries are created by grouping goroutines that share a common scheduler
state and trace (as determined by file locations) together.

## Install

```bash
$ go get github.com/felixge/gotraceanalyzer
```

## Usage

Analyze a traceback from a file:

```bash
$ gotraceanalyzer mypanic.txt
```

Or use stdin:

```bash
$ gotraceanalyzer < mypanic.txt
```

## Example

``` bash
$ curl -s https://raw.githubusercontent.com/felixge/gotraceanalyzer/master/example/output.txt | gotraceanalyzer
Showing 6 unique goroutines out of 230 total:

100 goroutines [sleep]:
time.Sleep(0x3b9aca00)
	/usr/local/go/src/pkg/runtime/time.goc:31 +0x31
main.Sleep()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:44 +0x26
main.D()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:40 +0x1a
main.C()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:36 +0x1a
main.B()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:32 +0x1a
main.A()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:28 +0x1a
created by main.main
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:9 +0x37

50 goroutines [sleep]:
time.Sleep(0x3b9aca00)
	/usr/local/go/src/pkg/runtime/time.goc:31 +0x31
main.Sleep()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:44 +0x26
main.D()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:40 +0x1a
main.C()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:36 +0x1a
main.B()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:32 +0x1a
created by main.main
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:11 +0x64

34 goroutines [sleep]:
time.Sleep(0x3b9aca00)
	/usr/local/go/src/pkg/runtime/time.goc:31 +0x31
main.Sleep()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:44 +0x26
main.D()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:40 +0x1a
main.C()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:36 +0x1a
created by main.main
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:14 +0xab

25 goroutines [sleep]:
time.Sleep(0x3b9aca00)
	/usr/local/go/src/pkg/runtime/time.goc:31 +0x31
main.Sleep()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:44 +0x26
main.D()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:40 +0x1a
created by main.main
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:17 +0xdc

20 goroutines [sleep]:
time.Sleep(0x3b9aca00)
	/usr/local/go/src/pkg/runtime/time.goc:31 +0x31
main.Sleep()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:44 +0x26
created by main.main
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:20 +0x122

1 goroutines [running]:
goroutine 1 [running]:
runtime.panic(0x26aa0, 0x2100e1050)
	/usr/local/go/src/pkg/runtime/panic.c:266 +0xb6
main.main()
	/Users/felix/code/go/src/github.com/felixge/gotraceanalyzer/example/main.go:24 +0x184
```

## TODO

* Support ignoring unrelated lines (useful if traceback got mixed up with log messages)

## License

MIT.
