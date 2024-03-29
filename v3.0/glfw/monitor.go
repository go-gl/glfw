package glfw

//#define GLFW_INCLUDE_NONE
//#include <GLFW/glfw3.h>
//GLFWmonitor* GetMonitorAtIndex(GLFWmonitor **monitors, int index);
//GLFWvidmode GetVidmodeAtIndex(GLFWvidmode *vidmodes, int index);
//void glfwSetMonitorCallbackCB();
//unsigned int GetGammaAtIndex(unsigned short *color, int i);
//void SetGammaAtIndex(unsigned short *color, int i, unsigned short value);
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

// GammaRamp describes the gamma ramp for a monitor.
type GammaRamp struct {
	Red   []uint16 //A slice of value describing the response of the red channel.
	Green []uint16 //A slice of value describing the response of the green channel.
	Blue  []uint16 //A slice of value describing the response of the blue channel.
}

// Monitor events.
const (
	Connected    MonitorEvent = C.GLFW_CONNECTED
	Disconnected MonitorEvent = C.GLFW_DISCONNECTED
)

// VideoMode describes a single video mode.
type VideoMode struct {
	Width       int //The width, in pixels, of the video mode.
	Height      int //The height, in pixels, of the video mode.
	RedBits     int //The bit depth of the red channel of the video mode.
	GreenBits   int //The bit depth of the green channel of the video mode.
	BlueBits    int //The bit depth of the blue channel of the video mode.
	RefreshRate int //The refresh rate, in Hz, of the video mode.
}

var fMonitorHolder func(monitor *Monitor, event MonitorEvent)

//export goMonitorCB
func goMonitorCB(monitor unsafe.Pointer, event C.int) {
	fMonitorHolder(&Monitor{(*C.GLFWmonitor)(monitor)}, MonitorEvent(event))
}

// GetMonitors returns a slice of handles for all currently connected monitors.
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

// GetPrimaryMonitor returns the primary monitor. This is usually the monitor
// where elements like the Windows task bar or the OS X menu bar is located.
func GetPrimaryMonitor() (*Monitor, error) {
	m := C.glfwGetPrimaryMonitor()

	if m == nil {
		return nil, errors.New("Can't get the primary monitor.")
	}
	return &Monitor{m}, nil
}

// GetPosition returns the position, in screen coordinates, of the upper-left
// corner of the monitor.
func (m *Monitor) GetPosition() (x, y int) {
	var xpos, ypos C.int

	C.glfwGetMonitorPos(m.data, &xpos, &ypos)
	return int(xpos), int(ypos)
}

// GetPhysicalSize returns the size, in millimetres, of the display area of the
// monitor.
//
// Note: Some operating systems do not provide accurate information, either
// because the monitor's EDID data is incorrect, or because the driver does not
// report it accurately.
func (m *Monitor) GetPhysicalSize() (width, height int) {
	var wi, h C.int

	C.glfwGetMonitorPhysicalSize(m.data, &wi, &h)
	return int(wi), int(h)
}

// GetName returns a human-readable name of the monitor, encoded as UTF-8.
func (m *Monitor) GetName() (string, error) {
	mn := C.glfwGetMonitorName(m.data)
	if mn == nil {
		return "", errors.New("Can't get monitor name.")
	}

	return C.GoString(mn), nil
}

// SetMonitorCallback sets the monitor configuration callback, or removes the
// currently set callback. This is called when a monitor is connected to or
// disconnected from the system.
func SetMonitorCallback(cbfun func(monitor *Monitor, event MonitorEvent)) {
	if cbfun == nil {
		C.glfwSetMonitorCallback(nil)
	} else {
		fMonitorHolder = cbfun
		C.glfwSetMonitorCallbackCB()
	}
}

// GetVideoModes returns an array of all video modes supported by the monitor.
// The returned array is sorted in ascending order, first by color bit depth
// (the sum of all channel depths) and then by resolution area (the product of
// width and height).
func (m *Monitor) GetVideoModes() ([]*VideoMode, error) {
	var length int

	vC := C.glfwGetVideoModes(m.data, (*C.int)(unsafe.Pointer(&length)))
	if vC == nil {
		return nil, errors.New("Can't get the video mode list.")
	}

	v := make([]*VideoMode, length)

	for i := 0; i < length; i++ {
		t := C.GetVidmodeAtIndex(vC, C.int(i))
		v[i] = &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits), int(t.refreshRate)}
	}

	return v, nil
}

// GetVideoMode returns the current video mode of the monitor. If you
// are using a full screen window, the return value will therefore depend on
// whether it is focused.
func (m *Monitor) GetVideoMode() (*VideoMode, error) {
	t := C.glfwGetVideoMode(m.data)

	if t == nil {
		return nil, errors.New("Can't get the video mode.")
	}
	return &VideoMode{int(t.width), int(t.height), int(t.redBits), int(t.greenBits), int(t.blueBits), int(t.refreshRate)}, nil
}

// SetGamma generates a 256-element gamma ramp from the specified exponent and then calls
// SetGamma with it.
func (m *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma(m.data, C.float(gamma))
}

// GetGammaRamp retrieves the current gamma ramp of the monitor.
func (m *Monitor) GetGammaRamp() (*GammaRamp, error) {
	var ramp GammaRamp

	rampC := C.glfwGetGammaRamp(m.data)
	if rampC == nil {
		return nil, errors.New("Can't get the gamma ramp.")
	}

	length := int(rampC.size)
	ramp.Red = make([]uint16, length)
	ramp.Green = make([]uint16, length)
	ramp.Blue = make([]uint16, length)

	for i := 0; i < length; i++ {
		ramp.Red[i] = uint16(C.GetGammaAtIndex(rampC.red, C.int(i)))
		ramp.Green[i] = uint16(C.GetGammaAtIndex(rampC.green, C.int(i)))
		ramp.Blue[i] = uint16(C.GetGammaAtIndex(rampC.blue, C.int(i)))
	}

	return &ramp, nil
}

// SetGammaRamp sets the current gamma ramp for the monitor.
func (m *Monitor) SetGammaRamp(ramp *GammaRamp) {
	var rampC C.GLFWgammaramp

	length := len(ramp.Red)

	for i := 0; i < length; i++ {
		C.SetGammaAtIndex(rampC.red, C.int(i), C.ushort(ramp.Red[i]))
		C.SetGammaAtIndex(rampC.green, C.int(i), C.ushort(ramp.Green[i]))
		C.SetGammaAtIndex(rampC.blue, C.int(i), C.ushort(ramp.Blue[i]))
	}

	C.glfwSetGammaRamp(m.data, &rampC)
}
