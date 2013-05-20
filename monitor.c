#include "_cgo_export.h"

GLFWmonitor** monitors;
const GLFWvidmode* vidmodes;

void InitMonitorArray(int* length) {
	monitors = glfwGetMonitors(length);
}

GLFWmonitor* GetMonitorAtIndex(int index) {
	return monitors[index];
}

void InitVidmodeArray(GLFWmonitor* monitor, int* length) {
	vidmodes = glfwGetVideoModes(monitor, length);
}

GLFWvidmode GetVidmodeAtIndex(int index) {
	return vidmodes[index];
}

void glfwSetMonitorCallbackCB() {
	glfwSetMonitorCallback(goMonitorCB)
}
