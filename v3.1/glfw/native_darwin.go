package glfw

//#define GLFW_EXPOSE_NATIVE_COCOA
//#define GLFW_EXPOSE_NATIVE_NSGL
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

// See: https://github.com/go-gl/glfw3/issues/82
/*
func (w *Window) GetCocoaWindow() C.id {
	ret := C.glfwGetCocoaWindow(w.data)
	panicError()
	return ret
}

func (w *Window) GetNSGLContext() C.id {
	ret := C.glfwGetNSGLContext(w.data)
	panicError()
	return ret
}
*/
