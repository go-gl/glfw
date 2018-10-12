package glfw

/*
#define GLFW_EXPOSE_NATIVE_COCOA
#define GLFW_EXPOSE_NATIVE_NSGL
#include "glfw/include/GLFW/glfw3.h"
#include "glfw/include/GLFW/glfw3native.h"

// workaround wrappers needed due to a cgo and/or LLVM bug.
// See: https://github.com/go-gl/glfw/issues/136
void *workaround_glfwGetCocoaWindow(GLFWwindow *w) {
	return (void *)glfwGetCocoaWindow(w);
}
void *workaround_glfwGetNSGLContext(GLFWwindow *w) {
	return (void *)glfwGetNSGLContext(w);
}
// workaround for inability to draw at start of program in OSX Mojave
// See: https://github.com/glfw/glfw/issues/1334
void cocoa_update_nsgl_context(void* id) {
    NSOpenGLContext *ctx = id;
    [ctx update];
}
*/
import "C"
import "unsafe"

// GetCocoaMonitor returns the CGDirectDisplayID of the monitor.
func (m *Monitor) GetCocoaMonitor() uintptr {
	ret := uintptr(C.glfwGetCocoaMonitor(m.data))
	panicError()
	return ret
}

// GetCocoaWindow returns the NSWindow of the window.
func (w *Window) GetCocoaWindow() uintptr {
	ret := uintptr(C.workaround_glfwGetCocoaWindow(w.data))
	panicError()
	return ret
}

// GetNSGLContext returns the NSOpenGLContext of the window.
func (w *Window) GetNSGLContext() uintptr {
	ret := uintptr(C.workaround_glfwGetNSGLContext(w.data))
	panicError()
	return ret
}

var numUpdates int

func (w *Window) updateNSGLContext() {
	if numUpdates < 2 {
		ctx := w.GetNSGLContext()
		C.cocoa_update_nsgl_context(unsafe.Pointer(ctx))
		numUpdates++
	}
}
