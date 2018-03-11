// Package bash wraps a patched version of bash that has hooks to call user-defined Go functions as shell builtins.
package bash

//go:generate bash build.sh 3.2.57

//#include "bash.h"
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	// mutex to guard the fns map
	mu  sync.Mutex
	fns = map[string]Function{}
)

func init() {
	C.go_init()
}

// Main starts bash with the given command line arguments (args) and environment variables (env).
// It returns the exit status code.
func Main(args, env []string) (status int) {
	c_argv := toCStringArray(args)
	c_env := toCStringArray(env)
	return int(C.Main(C.int(len(args)), c_argv, c_env))
}

// Register registers the fn function so it can be called from bash with the provided name.
// An error is returned if the name was already registered in this package.
func Register(name string, fn Function) error {
	mu.Lock()
	defer mu.Unlock()
	if fns[name] != nil {
		return fmt.Errorf("cgo-bash: function '%s' was already registered", name)
	}
	C.go_add_builtin(C.go_builtin{name: C.CString(name), function: (*C.sh_builtin_func_t)(C.builtin_func_wrapper)})
	fns[name] = fn
	return nil
}

// Unregister removes the function associated with the provided name.
func Unregister(name string) {
	mu.Lock()
	defer mu.Unlock()
	C.go_del_builtin(C.CString(name))
	delete(fns, name)
}

// Function represents a Go function that can be used as a bash builtin.
type Function func(args ...string) (status int)

//export builtin_func_wrapper
func builtin_func_wrapper(wl *C.WORD_LIST) C.int {
	args := make([]string, 0, 4)
	for ; wl != nil; wl = wl.next {
		args = append(args, C.GoString(wl.word.word))
	}
	fn := lookup(C.GoString(C.current_builtin.name))
	return C.int(fn(args...))
}

func lookup(name string) Function {
	mu.Lock()
	defer mu.Unlock()
	return fns[name]
}

func toCStringArray(a []string) **C.char {
	slice := make([]*C.char, len(a))
	for i, e := range a {
		slice[i] = C.CString(e)
	}
	data := (*reflect.SliceHeader)(unsafe.Pointer(&slice)).Data
	return (**C.char)(unsafe.Pointer(data))
}
