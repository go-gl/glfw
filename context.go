package glfw

//#include <stdlib.h>
//#include <GL/glfw3.h>
import "C"

import "unsafe"

func (w *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func GetCurrentContext() *Window {
	return (*Window)(unsafe.Pointer(C.glfwGetCurrentContext()))
}

func (w *Window) SwapBuffers() {
	C.glfwSwapBuffers((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func SwapInterval(interval int) {
	C.glfwSwapInterval(C.int(interval))
}

func ExtensionSupported(extension string) bool {
	w := C.CString(extension)
	defer C.free(unsafe.Pointer(w))
	r := C.glfwExtensionSupported(w)
	if r == C.GL_FALSE {
		return false
	}
	return true
}
