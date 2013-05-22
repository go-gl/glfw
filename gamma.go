package glfw

//#include <GL/glfw3.h>
import "C"

import "unsafe"

type GammaRamp C.GLFWgammaramp

//SetGamma generates a gamma ramp from the specified exponent and then calls
//SetGamma with it.
func (m *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma((*C.GLFWmonitor)(unsafe.Pointer(m)), C.float(gamma))
}

//GetGammaRamp retrieves the current gamma ramp of the specified monitor.
//
//NOTE: This function does not yet support monitors whose original gamma ramp
//has more or less than 256 entries.
func (m *Monitor) GetGammaRamp() *GammaRamp {
	ramp := new(C.GLFWgammaramp)
	C.glfwGetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), ramp)
	return (*GammaRamp)(unsafe.Pointer(ramp))
}

//SetGammaRamp sets the current gamma ramp for the specified monitor.
//
//NOTE: This function does not yet support monitors whose original gamma ramp
//has more or less than 256 entries.
func (m *Monitor) SetGammaRamp(ramp *GammaRamp) {
	C.glfwSetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), (*C.GLFWgammaramp)(unsafe.Pointer(ramp)))
}
