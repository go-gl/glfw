//+build linux freebsd

package glfw3

//#define GLFW_EXPOSE_NATIVE_X11
//#define GLFW_EXPOSE_NATIVE_GLX
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

func (w *Window) GetX11Window() C.Window {
	return C.glfwGetX11Window(w.data)
}

func (w *Window) GetGLXContext() C.GLXContext {
	return C.glfwGetGLXContext(w.data)
}

func GetX11Display() *C.Display {
	return C.glfwGetX11Display()
}
