# cgo-bash

## Install

```sh
$ go get -d github.com/tiborvass/cgo-bash
$ cd $GOPATH/src/github.com/tiborvass/cgo-bash
$ go generate # downloads, patches, and compiles bash
$ go install
```

## Usage

example.go:
```Go
package main

import (
	"fmt"
	bash "github.com/tiborvass/cgo-bash"
)

func init() {
	bash.Register("hello", hello)
}

func hello(args ...string) (status int) {
	fmt.Printf("Hello from Go! args=%v\n", args)
	return 0
}

func main() {
	os.Exit(bash.Main(os.Args, os.Environ()))
}
```

```sh
$ go run example.go
example-3.2# hello
Hello from Go! args=[]
example-3.2# hello world
Hello from Go! args=[world]
```

### Why?

Why not?


# License

MIT
