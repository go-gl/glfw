package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
import "C"

import "unsafe"

//SetClipboardString sets the system clipboard to the specified, UTF-8 encoded
//string.
func (w *Window) SetClipboardString(str string) {
	cp := C.CString(str)
	defer C.free(unsafe.Pointer(cp))
	C.glfwSetClipboardString((*C.GLFWwindow)(unsafe.Pointer(w)), cp)
}

//GetClipboardString returns the contents of the system clipboard, if it
//contains or is convertible to a UTF-8 encoded string.
func (w *Window) GetClipboardString() string {
	return C.GoString(C.glfwGetClipboardString((*C.GLFWwindow)(unsafe.Pointer(w))))
}
