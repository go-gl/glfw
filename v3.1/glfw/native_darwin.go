package glfw

/*
#define GLFW_EXPOSE_NATIVE_COCOA
#define GLFW_EXPOSE_NATIVE_NSGL
#include "glfw/include/GLFW/glfw3.h"
#include "glfw/include/GLFW/glfw3native.h"

// See: https://github.com/go-gl/glfw3/issues/82
inline void *workaround_glfwGetCocoaWindow(GLFWwindow *w) {
	return (void *)glfwGetCocoaWindow(w);
}
inline void *workaround_glfwGetNSGLContext(GLFWwindow *w) {
	return (void *)glfwGetNSGLContext(w);
}
*/
import "C"

func (w *Window) GetCocoaWindow() uintptr {
	ret := uintptr(C.workaround_glfwGetCocoaWindow(w.data))
	panicError()
	return ret
}

func (w *Window) GetNSGLContext() uintptr {
	ret := uintptr(C.workaround_glfwGetNSGLContext(w.data))
	panicError()
	return ret
}
