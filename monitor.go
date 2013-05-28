package glfw

//#include <GLFW/glfw3.h>
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

//GetMonitors returns a slice of handles for all currently connected monitors.
func GetMonitors() [](*Monitor) {
	var length int
	C.InitMonitorArray((*C.int)(unsafe.Pointer(&length)))
	m := make([](*Monitor), length)

	for i := 0; i < length; i++ {
		m[i] = (*Monitor)(unsafe.Pointer(C.GetMonitorAtIndex(C.int(i))))
	}

	return m
}

//GetPrimaryMonitor returns the primary monitor. This is usually the monitor
//where elements like the Windows task bar or the OS X menu bar is located.
func GetPrimaryMonitor() *Monitor {
	return (*Monitor)(unsafe.Pointer(C.glfwGetPrimaryMonitor()))
}

//GetPosition returns the position, in screen coordinates, of the upper-left
//corner of the specified monitor.
func (m *Monitor) GetPosition() (int, int) {
	var xpos, ypos C.int
	C.glfwGetMonitorPos((*C.GLFWmonitor)(unsafe.Pointer(m)), &xpos, &ypos)
	return int(xpos), int(ypos)
}

//GetPhysicalSize returns the size, in millimetres, of the display area of the
//specified monitor.
//
//Note: Some operating systems do not provide accurate information, either
//because the monitor's EDID data is incorrect, or because the driver does not
//report it accurately.
func (m *Monitor) GetPhysicalSize() (int, int) {
	var width, height C.int
	C.glfwGetMonitorPhysicalSize((*C.GLFWmonitor)(unsafe.Pointer(m)), &width, &height)
	return int(width), int(height)
}

//GetName returns a human-readable name, encoded as UTF-8, of the specified
//monitor.
func (m *Monitor) GetName() string {
	return C.GoString(C.glfwGetMonitorName((*C.GLFWmonitor)(unsafe.Pointer(m))))
}

//SetMonitorCallback sets the monitor configuration callback, or removes the
//currently set callback. This is called when a monitor is connected to or
//disconnected from the system.
//
//Function signature for this callback is: func(*Monitori int)
func SetMonitorCallback(cbfun goMonitorFunc) {
	fMonitorHolder = cbfun
	C.glfwSetMonitorCallbackCB()
}

//GetVideoModes returns an array of all video modes supported by the specified
//monitor. The returned array is sorted in ascending order, first by color bit
//depth (the sum of all channel depths) and then by resolution area (the product
//of width and height).
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

//GetVideoMode returns the current video mode of the specified monitor. If you
//are using a full screen window, the return value will therefore depend on
//whether it is focused.
func (m *Monitor) GetVideoMode() *VideoMode {
	t := C.glfwGetVideoMode((*C.GLFWmonitor)(unsafe.Pointer(m)))
	return &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits)}
}
