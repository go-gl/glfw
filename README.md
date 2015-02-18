Go Bindings for GLFW 3
======================

* This wrapper is now announced stable. There will be no API change until next major revision.
* API breaking changes are in *devel* branch.
* See [here](http://godoc.org/github.com/go-gl/glfw/3.0/glfw) for documentation.
* You can help by submitting examples to [go-gl/examples](http://github.com/go-gl/examples).

Remarks
=======

* Mingw64 users should rename ***glfw3dll.a*** to ***libglfw3dll.a***.
* In Windows and Linux, if you compile GLFW yourself, use <code>-DBUILD_SHARED_LIBS=on</code> with cmake in order to build the dynamic libraries.
* Some functions -which are marked in the documentation- can be called only from the main thread. Click [here](https://code.google.com/p/go-wiki/wiki/LockOSThread) for how.
* In OS X, you can install Go and GLFW via [Homebrew](http://brew.sh/).

```
$ brew install go
$ brew tap homebrew/versions
$ brew install --build-bottle --static glfw3
$ go get github.com/go-gl/glfw3
```

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
