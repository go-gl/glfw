package glfw3

//#define GLFW_EXPOSE_NATIVE_COCOA
//#define GLFW_EXPOSE_NATIVE_NSGL
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

// See: https://github.com/go-gl/glfw3/issues/82
/*
func (w *Window) GetCocoaWindow() (C.id, error) {
	return C.glfwGetCocoaWindow(w.data), fetchError()
}

func (w *Window) GetNSGLContext() (C.id, error) {
	return C.glfwGetNSGLContext(w.data), fetchError()
}
*/
