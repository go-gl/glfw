package glfw

//#include <GL/glfw3.h>
import "C"

import "unsafe"

type GammaRamp C.GLFWgammaramp

func (m *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma((*C.GLFWmonitor)(unsafe.Pointer(m)), C.float(gamma))
}

func (m *Monitor) GetGammaRamp() *GammaRamp {
	ramp := new(C.GLFWgammaramp)
	C.glfwGetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), ramp)
	return (*GammaRamp)(unsafe.Pointer(ramp))
}

func (m *Monitor) SetGammaRamp(ramp *GammaRamp) {
	C.glfwSetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), (*C.GLFWgammaramp)(unsafe.Pointer(ramp)))
}
