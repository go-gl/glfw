package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	if !glfw.Init() {
		panic(glfw.GetError())
	}
	defer glfw.Terminate()

	window := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if window == nil {
		panic(glfw.GetError())
	}

	window.MakeContextCurrent()

	window.SetCloseCallback(func(w *glfw.Window) {
		fmt.Println("hop")
	})

	window.SetCloseCallback(nil)
	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
