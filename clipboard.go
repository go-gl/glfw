package glfw3

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
import "C"

import (
	"errors"
	"unsafe"
)

//SetClipboardString sets the system clipboard to the specified UTF-8 encoded
//string.
func (w *Window) SetClipboardString(str string) {
	cp := C.CString(str)
	defer C.free(unsafe.Pointer(cp))

	// run in main thread
	do(func() {
		C.glfwSetClipboardString(w.data, cp)
	})
}

//GetClipboardString returns the contents of the system clipboard, if it
//contains or is convertible to a UTF-8 encoded string.
func (w *Window) GetClipboardString() (string, error) {
	var cs *C.char

	// run in main thread
	do(func() {
		cs = C.glfwGetClipboardString(w.data)
	})

	if cs == nil {
		return "", errors.New("Can't get clipboard string.")
	}

	return C.GoString(cs), nil
}
