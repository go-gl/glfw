#include "_cgo_export.h"

void glfwSetWindowPosCallbackCB(GLFWwindow *window) {
	glfwSetWindowPosCallback(window, goPositionCB)
}

void glfwSetWindowSizeCallbackCB(GLFWwindow *window) {
	glfwSetWindowSizeCallbackCB(window, goSizeCB)
}

void glfwSetWindowCloseCallbackCB(GLFWwindow *window) {
	glfwSetWindowCloseCallbackCB(window, goCloseCB)
}

void glfwSetWindowRefreshCallbackCB(GLFWwindow *window) {
	glfwSetWindowRefreshCallbackCB(window, goRefreshCB)
}

void glfwSetWindowFocusCallbackCB(GLFWwindow *window) {
	glfwSetWindowFocusCallbackCB(window, goFocusCB)
}

void glfwSetWindowIconifyCallbackCB(GLFWwindow *window) {
	glfwSetWindowIconifyCallbackCB(window, goIconifyCB)
}
