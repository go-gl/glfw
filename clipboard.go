package glfw3

//#include <stdlib.h>
//#include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"unsafe"
)

// SetClipboardString sets the system clipboard to the specified UTF-8 encoded
// string.
//
// This function may only be called from the main thread.
func (w *Window) SetClipboardString(str string) {
	cp := C.CString(str)
	defer C.free(unsafe.Pointer(cp))
	C.glfwSetClipboardString(w.data, cp)
}

// GetClipboardString returns the contents of the system clipboard, if it
// contains or is convertible to a UTF-8 encoded string.
//
// This function may only be called from the main thread.
func (w *Window) GetClipboardString() (string, error) {
	cs := C.glfwGetClipboardString(w.data)
	if cs == nil {
		return "", <-lastError
	}
	return C.GoString(cs), nil
}
