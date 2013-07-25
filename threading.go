package glfw3

import (
	"runtime"
)

var Inited bool

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()

	if !Init() {
		panic("Failed to initialize GLFW")
	}
	
	Inited = true
}

// Main runs the main glfw service loop.
// The binary's main.main must call glfw.Main() to run this loop.
// Main does not return. If the binary needs to do other work, it
// must do it in separate goroutines.
func Main() {
	for f := range mainfunc {
		f()
	}
}

// queue of work to run in main thread.
var mainfunc = make(chan func())

// do runs f on the main thread.
func do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}
