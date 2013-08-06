Go Bindings for GLFW 3
======================

* This wrapper is now announced stable. There will be no API change until next major revision.
* See http://godoc.org/github.com/go-gl/glfw3 for documentation.
* You can help by submitting examples to http://github.com/go-gl/examples

Remarks
=======

* Mingw64 users should rename ***glfw3dll.a*** to ***libglfw3dll.a***
* In Windows and Linux, if you compile GLFW yourself, use <code>-DBUILD_SHARED_LIBS=on</code> with cmake in order to build the dynamic libraries.
* Some functions -which are marked in the documentation- can be called only from the main thread. See https://code.google.com/p/go-wiki/wiki/LockOSThread for how.

Example
=======

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
