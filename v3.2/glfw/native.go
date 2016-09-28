package glfw

import "unsafe"

// GetGLFWWindow returns a *C.GLFWwindow reference (i.e. the GLFW window itself). This can be used for
// passing the GLFW window handle to external C libraries.
func (w *Window) GetGLFWWindow() uintptr {
	return uintptr(unsafe.Pointer(w.data))
}
