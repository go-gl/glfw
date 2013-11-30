Go Bindings for GLFW 3
======================

* **Warning:** This branch is for Go 1.2. See changelog for ***API breaking changes***.
* See [here](http://godoc.org/github.com/go-gl/glfw3) for documentation.
* You can help by submitting examples to [go-gl/examples](http://github.com/go-gl/examples).

Remarks
=======

* Mingw64 users should rename ***glfw3dll.a*** to ***libglfw3dll.a***.
* In Windows and Linux, if you compile GLFW yourself, use <code>-DBUILD_SHARED_LIBS=on</code> with cmake in order to build the dynamic libraries.
* Some functions -which are marked in the documentation- can be called only from the main thread. Click [here](https://code.google.com/p/go-wiki/wiki/LockOSThread) for how.
* In Mac OS, due to a bug in official Go packages, it's recommended to install Go and GLFW via [Homebrew](http://brew.sh/).

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
	glfw "github.com/go-gl/glfw3"
)

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
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

Changelog
=========

* <code>SetErrorCallback</code> This function is removed. The callback is now set internally. Functions return an error with corresponding code and description (do a type assertion to GlfwError for accessing the variables).
* <code>Init</code> Returns an error instead of bool.
* <code>GetTime</code> Returns an error.
* <code>GetCurrentContext</code> No longer returns an error.
* <code>GetJoystickAxes</code> No longer returns an error.
* <code>GetJoystickButtons</code> No longer returns an error.
* <code>GetJoystickName</code> No longer returns an error.
* <code>window.GetMonitor</code> No longer returns an error.
* <code>window.GetAttribute</code> Returns an error.
* <code>window.SetCharacterCallback</code> Accepts rune instead of uint.
