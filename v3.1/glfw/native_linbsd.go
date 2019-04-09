// +build linux freebsd

package glfw

/*
#define GLFW_EXPOSE_NATIVE_X11
#define GLFW_EXPOSE_NATIVE_GLX
#define GLFW_INCLUDE_NONE

#ifndef GO_GLFW_EXTERNAL
	#include "glfw/include/GLFW/glfw3.h"
	#include "glfw/include/GLFW/glfw3native.h"
#else
	#include <GLFW/glfw3.h>
	#include <GLFW/glfw3native.h>
#endif
*/
import "C"

func (w *Window) GetX11Window() C.Window {
	ret := C.glfwGetX11Window(w.data)
	panicError()
	return ret
}

func (w *Window) GetGLXContext() C.GLXContext {
	ret := C.glfwGetGLXContext(w.data)
	panicError()
	return ret
}

func GetX11Display() *C.Display {
	ret := C.glfwGetX11Display()
	panicError()
	return ret
}
