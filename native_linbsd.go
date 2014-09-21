//+build linux freebsd

package glfw3

//#define GLFW_EXPOSE_NATIVE_X11
//#define GLFW_EXPOSE_NATIVE_GLX
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

func (w *Window) GetX11Window() (C.Window, error) {
	return C.glfwGetX11Window(w.data), fetchError()
}

func (w *Window) GetGLXContext() (C.GLXContext, error) {
	return C.glfwGetGLXContext(w.data), fetchError()
}

func GetX11Display() (*C.Display, error) {
	return C.glfwGetX11Display(), fetchError()
}
