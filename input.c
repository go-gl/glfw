#include "_cgo_export.h"

void glfwSetKeyCallbackCB(GLFWwindow *window) {
	glfwSetKeyCallback(window, goKeyCB)
}

void glfwSetCharCallbackCB(GLFWwindow *window) {
	glfwSetCharCallback(window, goCharCB)
}

void glfwSetMouseCallbackCB(GLFWwindow *window) {
	glfwSetMouseButtonCallback(window, goMouseCB)
}

void glfwSetPosCallbackCB(GLFWwindow *window) {
	glfwSetCursorPosCallback(window, goPosCB)
}

void glfwSetEnterCallbackCB(GLFWwindow *window) {
	glfwSetCursorEnterCallback(window, goEnterCB)
}

void glfwSetScrollCallbackCB(GLFWwindow *window) {
	glfwSetScrollCallback(window, goScrollCB)
}
