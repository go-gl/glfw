#include "_cgo_export.h"

void glfwMouseCB(GLFWwindow* window, int button, int action, int mods) {
	goMouseCB(window, button, action, mods);
}

void glfwPosCB(GLFWwindow* window, double xpos, double ypos) {
	goPosCB(window, xpos, ypos);
}

void glfwEnterCB(GLFWwindow* window, int entered) {
	goEnterCB(window, entered);
}

void glfwScrollCB(GLFWwindow* window, double xpos, double ypos) {
	goScrollCB(window, xpos, ypos);
}

void glfwKeyCB(GLFWwindow* window, int key, int scancode, int action, int mods) {
	goKeyCB(window, key, scancode, action, mods);
}

void glfwCharCB(GLFWwindow* window, unsigned int character) {
	goCharCB(window, character);
}

void glfwSetKeyCallbackCB(GLFWwindow *window) {
	glfwSetKeyCallback(window, glfwKeyCB);
}

void glfwSetCharCallbackCB(GLFWwindow *window) {
	glfwSetCharCallback(window, glfwCharCB);
}

void glfwSetMouseCallbackCB(GLFWwindow *window) {
	glfwSetMouseButtonCallback(window, glfwMouseCB);
}

void glfwSetPosCallbackCB(GLFWwindow *window) {
	glfwSetCursorPosCallback(window, glfwPosCB);
}

void glfwSetEnterCallbackCB(GLFWwindow *window) {
	glfwSetCursorEnterCallback(window, glfwEnterCB);
}

void glfwSetScrollCallbackCB(GLFWwindow *window) {
	glfwSetScrollCallback(window, glfwScrollCB);
}

float GetAxisAtIndex(float *axis, int i) {
	return axis[i];
}

unsigned char GetButtonsAtIndex(unsigned char *buttons, int i) {
	return buttons[i];
}
