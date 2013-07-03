Go Bindings for GLFW 3
======================

* This wrapper is now announced stable. There will be no API change until next major revision.
* All the modules are implemented except "Native Acess".
* See http://godoc.org/github.com/go-gl/glfw3 for documentation.
* You can help by submitting examples to http://github.com/go-gl/examples

The library can be used as below:

```go
package main

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
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

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		//Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
```
