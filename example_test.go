package bash_test

import (
	"fmt"
	"os"

	bash "github.com/tiborvass/cgo-bash"
)

func ExampleRegister() {
	bash.Register("hello", hello)
	status := bash.Main([]string{os.Args[0], "-c", "hello world"}, os.Environ())
	fmt.Println("exit status", status)
	// Output:
	// Hello from Go! args=[world]
	// exit status 42
}

func hello(args ...string) (status int) {
	fmt.Printf("Hello from Go! args=%v\n", args)
	return 42
}
