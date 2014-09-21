package glfw3

//#define GLFW_EXPOSE_NATIVE_WIN32
//#define GLFW_EXPOSE_NATIVE_WGL
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

func (w *Window) GetWin32Window() (C.HWND, error) {
	return C.glfwGetWin32Window(w.data), fetchError()
}

func (w *Window) GetWGLContext() (C.HGLRC, error) {
	return C.glfwGetWGLContext(w.data), fetchError()
}
