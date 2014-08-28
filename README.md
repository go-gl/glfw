Go Bindings for GLFW 3
======================

* **ATTENTION:** As of GLFW 3.1 we break API. See Changelog below.
* See [here](http://godoc.org/github.com/go-gl/glfw3) for documentation.
* You can help by submitting examples to [go-gl/examples](http://github.com/go-gl/examples).

Remarks
=======

* Some functions -which are marked in the documentation- can be called only from the main thread. You need to use [runtime.LockOSThread()](http://godoc.org/runtime#LockOSThread) to arrange that main() runs on main thread.
* Installation is easy, just `go get github.com/go-gl/glfw3` and be done (*GLFW sources are included so you don't have to build GLFW on your own*)!

Example
=======

```go
package main

import (
	"runtime"

	glfw "github.com/go-gl/glfw3"
)

func init() {
	runtime.LockOSThread()
}

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
		// Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
```

Changelog
=========

* GLFW revision 5d525c4a5f9da0b8744f29affbf77d3a9580905c
* Easy `go get` installation (GLFW source code included in-repo and compiled in so you don't have to build GLFW on your own first and you don't have to distribute shared libraries).
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
* <code>window.SetDropCallback</code> added.
* <code>window.SetCharacterModsCallback</code> added.
* <code>PostEmptyEvent</code> added.
* Native window and context handlers added.
* Constant <code>ApiUnavailable</code> changed to <code>APIUnavailable</code>.
* Constant <code>ClientApi</code> changed to <code>ClientAPI</code>.
* Constant <code>OpenglForwardCompatible</code> changed to <code>OpenGLForwardCompatible</code>.
* Constant <code>OpenglDebugContext</code> changed to <code>OpenGLDebugContext</code>.
* Constant <code>OpenglProfile</code> changed to <code>OpenGLProfile</code>.
* Constant <code>SrgbCapable</code> changed to <code>SRGBCapable</code>.
* Constant <code>OpenglApi</code> changed to <code>OpenGLAPI</code>.
* Constant <code>OpenglEsApi</code> changed to <code>OpenGLESAPI</code>.
* Constant <code>OpenglAnyProfile</code> changed to <code>OpenGLAnyProfile</code>.
* Constant <code>OpenglCoreProfile</code> changed to <code>OpenGLCoreProfile</code>.
* Constant <code>OpenglCompatProfile</code> changed to <code>OpenGLCompatProfile</code>.
