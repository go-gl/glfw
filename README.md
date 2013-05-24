GLFW 3.0 (WIP) Bindings for Go
==============================

ATTENTION: GLFW 3.0 is still in development and the API changes occur very often. Current wrapper is built with master elmindreda/glfw checked out at 20.05.2013.

* All the modules are implemented except "Native Acess".
* See http://godoc.org/github.com/tapir/glfw3-go for documentation.
* You can help by submitting examples.

The library can be used as below:
<pre>
	package main
	
	import (
		"fmt"
		glfw "github.com/tapir/glfw3-go"
	)
	
	func errorCallback(err int, desc string) {
		fmt.Printf("%v: %v\n", err, desc)
	}
	
	func main() {
		glfw.SetErrorCallback(errorCallback)
	
		if !glfw.Init() {
			panic("Can't init glfw!")
		}
		defer glfw.Terminate()
	
		window := glfw.CreateWindow(640, 480, "Testing", nil, nil)
		if window == nil {
			panic("Can't create window!")
		}
		defer window.Destroy()
	
		window.MakeContextCurrent()
	
		for {
			if window.ShouldClose() {
				break
			}
		
			glfw.PollEvents()
		}
	}
</pre>
