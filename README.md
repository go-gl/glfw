GLFW 3.0 Bindings for Go
========================

* All the modules are implemented except "Native Acess".
* See http://godoc.org/github.com/tapir/glfw3-go for documentation.
* You can help by submitting examples to http://github.com/tapir/glfw3-go-examples

The library can be used as below:
<pre>
	package main
	
	import (
		"fmt"
		glfw "github.com/tapir/glfw3-go"
	)
	
	func errorCallback(err glfw.ErrorCode, desc string) {
		fmt.Printf("%v: %v\n", err, desc)
	}
	
	func main() {
		glfw.SetErrorCallback(errorCallback)
	
		if !glfw.Init() {
			panic("Can't init glfw!")
		}
		defer glfw.Terminate()
	
		window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
		if err != nil {
			panic(err)
		}
		defer window.Destroy()
	
		window.MakeContextCurrent()
	
		for !window.ShouldClose() {
			//Do OpenGL stuff
			glfw.PollEvents()
		}
	}
</pre>
