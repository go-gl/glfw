// +build linux freebsd

package glfw

//#define GLFW_EXPOSE_NATIVE_X11
//#define GLFW_EXPOSE_NATIVE_GLX
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

func GetX11Display() *C.Display {
	ret := C.glfwGetX11Display()
	panicError()
	return ret
}

func (m *Monitor) GetX11Adapter() C.RRCrtc {
	ret := C.glfwGetX11Adapter(m.data)
	panicError()
	return ret
}

func (m *Monitor) GetX11Monitor() C.RROutput {
	ret := C.glfwGetX11Monitor(m.data)
	panicError()
	return ret
}

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

func (w *Window) GetGLXWindow() C.GLXWindow {
	ret := C.glfwGetGLXWindow(w.data)
	panicError()
	return ret
}
