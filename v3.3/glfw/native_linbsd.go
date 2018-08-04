// +build linux,!wayland freebsd,!wayland

package glfw

//#define GLFW_EXPOSE_NATIVE_X11
//#define GLFW_EXPOSE_NATIVE_GLX
//#include "glfw/include/GLFW/glfw3.h"
//#include "glfw/include/GLFW/glfw3native.h"
import "C"

// GetX11Display returns X11 display handle
func GetX11Display() *C.Display {
	return C.glfwGetX11Display()
}

// GetX11Adapter returns the RRCrtc of the monitor.
func (m *Monitor) GetX11Adapter() C.RRCrtc {
	return C.glfwGetX11Adapter(m.data)
}

// GetX11Monitor returns the RROutput of the monitor.
func (m *Monitor) GetX11Monitor() C.RROutput {
	return C.glfwGetX11Monitor(m.data)
}

// GetX11Window returns the Window of the window.
func (w *Window) GetX11Window() C.Window {
	return C.glfwGetX11Window(w.data)
}

// GetGLXContext returns the GLXContext of the window.
func (w *Window) GetGLXContext() C.GLXContext {
	return C.glfwGetGLXContext(w.data)
}

// GetGLXWindow returns the GLXWindow of the window.
func (w *Window) GetGLXWindow() C.GLXWindow {
	return C.glfwGetGLXWindow(w.data)
}
