package glfw

//#define NS_FORMAT_ARGUMENT(A)
//#define GLFW_EXPOSE_NATIVE_WIN32
//#define GLFW_EXPOSE_NATIVE_WGL
//#define GLFW_INCLUDE_NONE
//#include <GLFW/glfw3.h>
//#include <GLFW/glfw3native.h>
import "C"

func (w *Window) GetWin32Window() C.HWND {
	return C.glfwGetWin32Window(w.data)
}

func (w *Window) GetWGLContext() C.HGLRC {
	return C.glfwGetWGLContext(w.data)
}
