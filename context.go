package glfw3

//#include <stdlib.h>
//#include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"unsafe"
)

// MakeContextCurrent makes the context of the window current.
// Originally GLFW 3 passes a null pointer to detach the context.
// But since we're using receievers, DetachCurrentContext should
// be used instead.
func (w *Window) MakeContextCurrent() error {
	C.glfwMakeContextCurrent(w.data)
	return fetchError()
}

// DetachCurrentContext detaches the current context.
func DetachCurrentContext() error {
	C.glfwMakeContextCurrent(nil)
	return fetchError()
}

// GetCurrentContext returns the window whose context is current.
func GetCurrentContext() (*Window, error) {
	w := C.glfwGetCurrentContext()
	if w == nil {
		return nil, fetchError()
	}
	return windows.get(w), nil
}

// SwapBuffers swaps the front and back buffers of the window. If the
// swap interval is greater than zero, the GPU driver waits the specified number
// of screen updates before swapping the buffers.
func (w *Window) SwapBuffers() error {
	C.glfwSwapBuffers(w.data)
	return fetchError()
}

// SwapInterval sets the swap interval for the current context, i.e. the number
// of screen updates to wait before swapping the buffers of a window and
// returning from SwapBuffers. This is sometimes called
// 'vertical synchronization', 'vertical retrace synchronization' or 'vsync'.
//
// Contexts that support either of the WGL_EXT_swap_control_tear and
// GLX_EXT_swap_control_tear extensions also accept negative swap intervals,
// which allow the driver to swap even if a frame arrives a little bit late.
// You can check for the presence of these extensions using
// ExtensionSupported. For more information about swap tearing,
// see the extension specifications.
//
// Some GPU drivers do not honor the requested swap interval, either because of
// user settings that override the request or due to bugs in the driver.
func SwapInterval(interval int) error {
	C.glfwSwapInterval(C.int(interval))
	return fetchError()
}

// ExtensionSupported returns whether the specified OpenGL or context creation
// API extension is supported by the current context. For example, on Windows
// both the OpenGL and WGL extension strings are checked.
//
// As this functions searches one or more extension strings on each call, it is
// recommended that you cache its results if it's going to be used frequently.
// The extension strings will not change during the lifetime of a context, so
// there is no danger in doing this.
func ExtensionSupported(extension string) (bool, error) {
	e := C.CString(extension)
	defer C.free(unsafe.Pointer(e))
	return glfwbool(C.glfwExtensionSupported(e)), fetchError()
}
