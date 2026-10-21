//go:build (linux || freebsd || netbsd || openbsd) && ((!x11 && !wayland) || wayland)

package glfw

/*
#include "glfw/src/internal.h"

static GLFWmonitor* waylandWindowMonitor(GLFWwindow* handle)
{
#if defined(_GLFW_WAYLAND)
	_GLFWwindow* window = (_GLFWwindow*) handle;
	if (_glfw.platform.platformID != GLFW_PLATFORM_WAYLAND || window->wl.outputScaleCount == 0)
		return NULL;

	struct wl_output* output = window->wl.outputScales[window->wl.outputScaleCount - 1].output;
	for (int i = 0; i < _glfw.monitorCount; i++)
	{
		if (_glfw.monitors[i]->wl.output == output)
			return (GLFWmonitor*) _glfw.monitors[i];
	}
#endif
	return NULL;
}
*/
import "C"

// GetWaylandWindowMonitor returns the monitor of an output the window's
// surface is on, as Wayland gives windows no position to find it by. When the
// surface is on two outputs it is the one entered last. It returns nil when the
// surface is on no output or GLFW is not running on Wayland.
//
// This function must only be called from the main thread.
func (w *Window) GetWaylandWindowMonitor() *Monitor {
	m := C.waylandWindowMonitor(w.data)
	if m == nil {
		return nil
	}
	return &Monitor{m}
}
