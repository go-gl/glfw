GLFW 3.0 (WIP) Bindings for Go
==============================

**ATTENTION:** GLFW 3.0 is still in development and the API changes occur very often. I try to stay updated with the latest changes on https://github.com/glfw/glfw _(Last update: c159411944c68c79797321bc76529a351544236a)_

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
	
		window := glfw.CreateWindow(640, 480, "Testing", &glfw.Monitor{}, &glfw.Window{})
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
