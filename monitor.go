package glfw

//#include <GL/glfw3.h>
//void InitMonitorArray(int* length);
//GLFWmonitor* GetMonitorAtIndex(int index);
//void InitVidmodeArray(GLFWmonitor* monitor, int* length);
//GLFWvidmode GetVidmodeAtIndex(int index);
//void glfwSetMonitorCallbackCB();
import "C"

import "unsafe"

type Monitor C.GLFWmonitor

type VideoMode struct {
	Width     int
	Height    int
	RedBits   int
	GreenBits int
	BlueBits  int
}

type goMonitorFunc func(*Monitor, int)

var fMonitorHolder goMonitorFunc

//export goMonitorCB
func goMonitorCB(monitor unsafe.Pointer, event C.int) {
	fMonitorHolder((*Monitor)(unsafe.Pointer(monitor)), int(event))
}

func GetMonitors() [](*Monitor) {
	var length int
	C.InitMonitorArray((*C.int)(unsafe.Pointer(&length)))
	m := make([](*Monitor), length)

	for i := 0; i < length; i++ {
		m[i] = (*Monitor)(unsafe.Pointer(C.GetMonitorAtIndex(C.int(i))))
	}

	return m
}

func GetPrimaryMonitor() *Monitor {
	return (*Monitor)(unsafe.Pointer(C.glfwGetPrimaryMonitor()))
}

func (m *Monitor) GetPosition() (int, int) {
	var xpos, ypos C.int
	C.glfwGetMonitorPos((*C.GLFWmonitor)(unsafe.Pointer(m)), &xpos, &ypos)
	return int(xpos), int(ypos)
}

func (m *Monitor) GetPhysicalSize() (int, int) {
	var width, height C.int
	C.glfwGetMonitorPhysicalSize((*C.GLFWmonitor)(unsafe.Pointer(m)), &width, &height)
	return int(width), int(height)
}

func (m *Monitor) GetName() string {
	return C.GoString(C.glfwGetMonitorName((*C.GLFWmonitor)(unsafe.Pointer(m))))
}

func SetMonitorCallback(cbfun goMonitorFunc) {
	fMonitorHolder = cbfun
	C.glfwSetMonitorCallbackCB()
}

func (m *Monitor) GetVideoModes() [](*VideoMode) {
	var length int
	C.InitVidmodeArray((*C.GLFWmonitor)(unsafe.Pointer(m)), (*C.int)(unsafe.Pointer(&length)))
	v := make([](*VideoMode), length)

	for i := 0; i < length; i++ {
		t := C.GetVidmodeAtIndex(C.int(i))
		v[i] = &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits)}
	}

	return v
}

func (m *Monitor) GetVideoMode() *VideoMode {
	t := C.glfwGetVideoMode((*C.GLFWmonitor)(unsafe.Pointer(m)))
	return &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits)}
}
