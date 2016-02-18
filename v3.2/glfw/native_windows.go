package glfw

//#define GLFW_EXPOSE_NATIVE_WIN32
//#define GLFW_EXPOSE_NATIVE_WGL
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

func (m *Monitor) GetWin32Adapter() string {
	ret := C.glfwGetWin32Adapter(m.data)
	panicError()
	return C.GoString(ret)
}

func (m *Monitor) GetWin32Monitor() string {
	ret := C.glfwGetWin32Monitor(m.data)
	panicError()
	return C.GoString(ret)
}

func (w *Window) GetWin32Window() C.HWND {
	ret := C.glfwGetWin32Window(w.data)
	panicError()
	return ret
}

func (w *Window) GetWGLContext() C.HGLRC {
	ret := C.glfwGetWGLContext(w.data)
	panicError()
	return ret
}
