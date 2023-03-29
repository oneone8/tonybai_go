// trace.go
package samps

import "testing"
import "time"

func Trace(name string) func() {
	println("enter:", name)
	return func() {
		println("exit:", name)
	}
}

func foo() {
	defer Trace("foo")()
	bar()
}

func bar() {
	defer Trace("bar")()
}

func Test5(t *testing.T) {
	defer Trace("main")()

	time.Sleep(5 * time.Second)
	foo()
}
