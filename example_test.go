package bash_test

import (
	"fmt"
	"os"

	bash "github.com/tiborvass/cgo-bash"
)

// ExampleRegister shows how you can use bash.Main
func ExampleRegister() {
	bash.Register("hello", Hello)
	status := bash.Main([]string{os.Args[0], "-c", "hello world"}, os.Environ())
	fmt.Println("exit status", status)
	// Output:
	// Hello from Go! args=[world]
	// exit status 42
}

func Hello(args ...string) (status int) {
	fmt.Printf("Hello from Go! args=%v\n", args)
	return 42
}
