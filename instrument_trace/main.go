package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 1. 使用defer函数跟踪函数的执行过程
// 2. 不需要手动传入函数的名字
// 3. 需要打印协程id
// 4. 需要有层次感
// 5. 可以自动添加defer Trace()()

var goroutineSpace = []byte("goroutine ")
var m = &sync.Map{}

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(gid uint64, name, arrow string, indent int) {
	indents := ""

	for i := 0; i < indent; i++ {
		// 这里可以优化速度哦
		indents += "	"
	}

	fmt.Printf("g[%05d]: %s%s%s\n", gid, indents, arrow, name)
}

func Trace() func() {

	// https://www.cnblogs.com/aganippe/p/16285871.html
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		panic("not found caller")
	}

	callerName := runtime.FuncForPC(pc).Name()

	gid := curGoroutineID()

	indent, ok := m.Load(gid)
	if !ok {
		indent = 0
	}
	m.Store(gid, indent.(int)+1)

	printTrace(gid, callerName, "->", indent.(int))

	return func() {
		indent, ok := m.Load(gid)
		if !ok {
			panic("not found gid")
		} else {
			m.Store(gid, indent.(int)-1)
		}
		printTrace(gid, callerName, "<-", indent.(int)-1)
	}

}

func foo() {
	defer Trace()()
}

func zoo() {
	defer Trace()()
	foo()
}

func main() {
	defer Trace()()
	zoo()
}
