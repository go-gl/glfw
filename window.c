#include "_cgo_export.h"

void glfwPositionCB(GLFWwindow* window, int xpos, int ypos) {
	goPositionCB(window, xpos, ypos);
}

void glfwSizeCB(GLFWwindow* window, int width, int height) {
	goSizeCB(window, width, height);
}

void glfwCloseCB(GLFWwindow* window) {
	goCloseCB(window);
}

void glfwRefreshCB(GLFWwindow* window) {
	goRefreshCB(window);
}

void glfwFocusCB(GLFWwindow* window, int focused) {
	goFocusCB(window, focused);
}

void glfwIconifyCB(GLFWwindow* window, int iconified) {
	goIconifyCB(window, iconified);
}

void glfwSetWindowPosCallbackCB(GLFWwindow* window) {
	glfwSetWindowPosCallback(window, glfwPositionCB);
}

void glfwSetWindowSizeCallbackCB(GLFWwindow* window) {
	glfwSetWindowSizeCallback(window, glfwSizeCB);
}

void glfwSetWindowCloseCallbackCB(GLFWwindow* window) {
	glfwSetWindowCloseCallback(window, glfwCloseCB);
}

void glfwSetWindowRefreshCallbackCB(GLFWwindow* window) {
	glfwSetWindowRefreshCallback(window, glfwRefreshCB);
}

void glfwSetWindowFocusCallbackCB(GLFWwindow* window) {
	glfwSetWindowFocusCallback(window, glfwFocusCB);
}

void glfwSetWindowIconifyCallbackCB(GLFWwindow* window) {
	glfwSetWindowIconifyCallback(window, glfwIconifyCB);
}
