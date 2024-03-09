package main

import (
	"image"
	"image/color"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
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

	window, err := glfw.CreateWindow(640, 480, "Custom Cursor", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Creating a custom cursor.
	cursor := glfw.CreateCursor(whiteTriangle, 0, 0)
	window.SetCursor(cursor)

	// Setting a custom cursor.
	window.SetIcon([]image.Image{whiteTriangle})

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var whiteTriangle = func() *image.NRGBA {
	c := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	const size = 16
	m := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size-y; x++ {
			m.SetNRGBA(x, y, c)
		}
	}
	return m
}()
