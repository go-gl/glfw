# Go Bindings for GLFW 3.1

* **ATTENTION:** As of GLFW 3.1 we break API. See Changelog below.
* See [here](http://godoc.org/github.com/go-gl/glfw/3.1/glfw) for documentation.
* You can help by submitting examples to [go-gl/examples](http://github.com/go-gl/examples).

## Installation

* Installation is easy, just `go get github.com/go-gl/glfw3` and be done (*GLFW sources are included so you don't have to build GLFW on your own*)!
* Go 1.4 is required on Windows (otherwise you must use MinGW v4.8.1 exactly, see [Go issue 8811](https://code.google.com/p/go/issues/detail?id=8811)).

## Usage

```go
package main

import (
	"runtime"
	glfw "github.com/go-gl/glfw3"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
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

## Changelog

* Easy `go get` installation. GLFW source code is now included in-repo and compiled in so you don't have to build GLFW on your own and distribute shared libraries. The revision of GLFW C library used is listed in [GLFW_C_REVISION.txt](GLFW_C_REVISION.txt) file.
* The error callback is now set internally. Functions return an error with corresponding code and description (do a type assertion to glfw3.Error for accessing the variables) if the error is recoverable. If not a panic will occur.

### GLFW 3.1 Specfic Changes
* Added type `Cursor`.
* Added function `Window.SetDropCallback`.
* Added function `Window.SetCharModsCallback`.
* Added function `PostEmptyEvent`.
* Added function `CreateCursor`.
* Added function `CreateStandardCursor`.
* Added function `Cursor.Destroy`.
* Added function `Window.SetCursor`.
* Added function `Window.GetFrameSize`.
* Added window hint `Floating`.
* Added window hint `AutoIconify`.
* Added window hint `ContextReleaseBehavior`.
* Added window hint `DoubleBuffer`.
* Added hint value `AnyReleaseBehavior`.
* Added hint value `ReleaseBehaviorFlush`.
* Added hint value `ReleaseBehaviorNone`.
* Added hint value `DontCare`.

### API changes
* `Window.Iconfiy` Returns an error.
* `Window.Restore` Returns an error.
* `Init` Returns an error instead of `bool`.
* `GetJoystickAxes` No longer returns an error.
* `GetJoystickButtons` No longer returns an error.
* `GetJoystickName` No longer returns an error.
* `GetMonitors` No longer returns an error.
* `GetPrimaryMonitor` No longer returns an error.
* `Monitor.GetGammaRamp` No longer returns an error.
* `Monitor.GetVideoMode` No longer returns an error.
* `Monitor.GetVideoModes` No longer returns an error.
* `GetCurrentContext` No longer returns an error.
* `Window.SetCharCallback` Accepts `rune` instead of `uint`.
* Added type `Error`.
* Removed `SetErrorCallback`.
* Removed error code `NotInitialized`.
* Removed error code `NoCurrentContext`.
* Removed error code `InvalidEnum`.
* Removed error code `InvalidValue`.
* Removed error code `OutOfMemory`.
* Removed error code `PlatformError`.
* Removed `KeyBracket`.
* Renamed `Window.SetCharacterCallback` to `Window.SetCharCallback`.
* Renamed `Window.GetCursorPosition` to `GetCursorPos`.
* Renamed `Window.SetCursorPosition` to `SetCursorPos`.
* Renamed `CursorPositionCallback` to `CursorPosCallback`.
* Renamed `Window.SetCursorPositionCallback` to `SetCursorPosCallback`.
* Renamed `VideoMode` to `VidMode`.
* Renamed `Monitor.GetPosition` to `Monitor.GetPos`.
* Renamed `Window.GetPosition` to `Window.GetPos`.
* Renamed `Window.SetPosition` to `Window.SetPos`.
* Renamed `Window.GetAttribute` to `Window.GetAttrib`.
* Renamed `Window.SetPositionCallback` to `Window.SetPosCallback`.
* Renamed `PositionCallback` to `PosCallback`.
* Ranamed `Cursor` to `CursorMode`.
* Renamed `StickyKeys` to `StickyKeysMode`.
* Renamed `StickyMouseButtons` to `StickyMouseButtonsMode`.
* Renamed `ApiUnavailable` to `APIUnavailable`.
* Renamed `ClientApi` to `ClientAPI`.
* Renamed `OpenglForwardCompatible` to `OpenGLForwardCompatible`.
* Renamed `OpenglDebugContext` to `OpenGLDebugContext`.
* Renamed `OpenglProfile` to `OpenGLProfile`.
* Renamed `SrgbCapable` to `SRGBCapable`.
* Renamed `OpenglApi` to `OpenGLAPI`.
* Renamed `OpenglEsApi` to `OpenGLESAPI`.
* Renamed `OpenglAnyProfile` to `OpenGLAnyProfile`.
* Renamed `OpenglCoreProfile` to `OpenGLCoreProfile`.
* Renamed `OpenglCompatProfile` to `OpenGLCompatProfile`.
* Renamed `KeyKp...` to `KeyKP...`.
