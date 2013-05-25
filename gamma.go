package glfw

//#include <GLFW/glfw3.h>
//unsigned int GetGammaAtIndex(unsigned short *color, int i);
//void SetGammaAtIndex(unsigned short *color, int i, unsigned short value);
import "C"

import "unsafe"

//GammaRamp describes the gamma ramp for a monitor.
type GammaRamp struct {
	Red   []uint16
	Green []uint16
	Blue  []uint16
	Size  uint
}

//SetGamma generates a gamma ramp from the specified exponent and then calls
//SetGamma with it.
func (m *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma((*C.GLFWmonitor)(unsafe.Pointer(m)), C.float(gamma))
}

//GetGammaRamp retrieves the current gamma ramp of the specified monitor.
func (m *Monitor) GetGammaRamp() *GammaRamp {
	var ramp GammaRamp

	rampC := C.glfwGetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)))
	length := int(rampC.size)

	ramp.Red = make([]uint16, length)
	ramp.Green = make([]uint16, length)
	ramp.Blue = make([]uint16, length)
	ramp.Size = uint(length)

	for i := 0; i < length; i++ {
		ramp.Red[i] = uint16(C.GetGammaAtIndex(rampC.red, C.int(i)))
		ramp.Green[i] = uint16(C.GetGammaAtIndex(rampC.green, C.int(i)))
		ramp.Blue[i] = uint16(C.GetGammaAtIndex(rampC.blue, C.int(i)))
	}

	return &ramp
}

//SetGammaRamp sets the current gamma ramp for the specified monitor.
func (m *Monitor) SetGammaRamp(ramp *GammaRamp) {
	var rampC C.GLFWgammaramp

	length := int(ramp.Size)

	for i := 0; i < length; i++ {
		C.SetGammaAtIndex(rampC.red, C.int(i), C.ushort(ramp.Red[i]))
		C.SetGammaAtIndex(rampC.green, C.int(i), C.ushort(ramp.Green[i]))
		C.SetGammaAtIndex(rampC.blue, C.int(i), C.ushort(ramp.Blue[i]))
	}
	rampC.size = C.uint(ramp.Size)

	C.glfwSetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), &rampC)
}
