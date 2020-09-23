# Hello, Go compiler!

Some hello-world ways to debug Go programs and compiler.

* Dump the SSAs

```bash
$ env GOSSAFUNC=main go build main.go
```

* Dump the assembly

```bash
$ go build -gcflags -S main.go
```

(The `gc` here means *Go compiler*, that's really some masterpiece of wording)
