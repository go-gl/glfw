package glfw

//#include <GLFW/glfw3.h>
//GLFWmonitor* GetMonitorAtIndex(GLFWmonitor **monitors, int index);
//GLFWvidmode GetVidmodeAtIndex(GLFWvidmode *vidmodes, int index);
//void glfwSetMonitorCallbackCB();
import "C"

import (
	"errors"
	"unsafe"
)

type Monitor struct {
	data *C.GLFWmonitor
}

// MonitorEvent corresponds to a monitor configuration event.
type MonitorEvent int

// Monitor events.
const (
	Connected    MonitorEvent = C.GLFW_CONNECTED
	Disconnected MonitorEvent = C.GLFW_DISCONNECTED
)

//VideoMode describes a single video mode.
type VideoMode struct {
	Width     int //The width, in pixels, of the video mode.
	Height    int //The height, in pixels, of the video mode.
	RedBits   int //The bit depth of the red channel of the video mode.
	GreenBits int //The bit depth of the green channel of the video mode.
	BlueBits  int //The bit depth of the blue channel of the video mode.
}

var fMonitorHolder func(monitor *Monitor, event MonitorEvent)

//export goMonitorCB
func goMonitorCB(monitor unsafe.Pointer, event C.int) {
	fMonitorHolder(&Monitor{(*C.GLFWmonitor)(unsafe.Pointer(monitor))}, MonitorEvent(event))
}

//GetMonitors returns a slice of handles for all currently connected monitors.
func GetMonitors() ([]*Monitor, error) {
	var length int

	mC := C.glfwGetMonitors((*C.int)(unsafe.Pointer(&length)))

	if mC == nil {
		return nil, errors.New("Can't get the monitor list.")
	}

	m := make([]*Monitor, length)

	for i := 0; i < length; i++ {
		m[i] = &Monitor{C.GetMonitorAtIndex(mC, C.int(i))}
	}

	return m, nil
}

//GetPrimaryMonitor returns the primary monitor. This is usually the monitor
//where elements like the Windows task bar or the OS X menu bar is located.
func GetPrimaryMonitor() (*Monitor, error) {
	m := C.glfwGetPrimaryMonitor()

	if m == nil {
		return nil, errors.New("Can't get the primary monitor.")
	}
	return &Monitor{m}, nil
}

//GetPosition returns the position, in screen coordinates, of the upper-left
//corner of the monitor.
func (m *Monitor) GetPosition() (int, int) {
	var xpos, ypos C.int

	C.glfwGetMonitorPos(m.data, &xpos, &ypos)
	return int(xpos), int(ypos)
}

//GetPhysicalSize returns the size, in millimetres, of the display area of the
//monitor.
//
//Note: Some operating systems do not provide accurate information, either
//because the monitor's EDID data is incorrect, or because the driver does not
//report it accurately.
func (m *Monitor) GetPhysicalSize() (int, int) {
	var width, height C.int

	C.glfwGetMonitorPhysicalSize(m.data, &width, &height)
	return int(width), int(height)
}

//GetName returns a human-readable name of the monitor, encoded as UTF-8.
func (m *Monitor) GetName() string {
	return C.GoString(C.glfwGetMonitorName(m.data))
}

//SetMonitorCallback sets the monitor configuration callback, or removes the
//currently set callback. This is called when a monitor is connected to or
//disconnected from the system.
func SetMonitorCallback(cbfun func(monitor *Monitor, event MonitorEvent)) {
	fMonitorHolder = cbfun
	C.glfwSetMonitorCallbackCB()
}

//GetVideoModes returns an array of all video modes supported by the monitor.
//The returned array is sorted in ascending order, first by color bit depth
//(the sum of all channel depths) and then by resolution area (the product of
//width and height).
func (m *Monitor) GetVideoModes() ([]*VideoMode, error) {
	var length int

	vC := C.glfwGetVideoModes(m.data, (*C.int)(unsafe.Pointer(&length)))
	if vC == nil {
		return nil, errors.New("Can't get the video mode list.")
	}

	v := make([]*VideoMode, length)

	for i := 0; i < length; i++ {
		t := C.GetVidmodeAtIndex(vC, C.int(i))
		v[i] = &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits)}
	}

	return v, nil
}

//GetVideoMode returns the current video mode of the monitor. If you
//are using a full screen window, the return value will therefore depend on
//whether it is focused.
func (m *Monitor) GetVideoMode() (*VideoMode, error) {
	t := C.glfwGetVideoMode(m.data)

	if t == nil {
		return nil, errors.New("Can't get the video mode.")
	}
	return &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits)}, nil
}
