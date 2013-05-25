#include "_cgo_export.h"

void glfwErrorCB(int err, const char *desc) {
	goErrorCB(err, (char*)desc);
}

void glfwSetErrorCallbackCB() {
	glfwSetErrorCallback(glfwErrorCB);
}