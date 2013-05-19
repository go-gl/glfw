package glfw

//#include <stdlib.h>
//#include <GL/glfw3.h>
import "C"

import "unsafe"

func (w *Window) SetClipboardString(str string) {
	cp := C.CString(str)
	defer C.free(unsafe.Pointer(cp))
	C.glfwSetClipboardString((*C.GLFWwindow)(unsafe.Pointer(w)), cp)
}

func (w *Window) GetClipboardString() string {
	return C.GoString(C.glfwGetClipboardString((*C.GLFWwindow)(unsafe.Pointer(w))))
}
