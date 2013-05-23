package glfw

//#include <GL/glfw3.h>
import "C"

import "unsafe"

const GammaRampSize = C.GLFW_GAMMA_RAMP_SIZE

type GammaRamp struct {
	Red [GammaRampSize]uint16
	Green [GammaRampSize]uint16
	Blue [GammaRampSize]uint16
}

//SetGamma generates a gamma ramp from the specified exponent and then calls
//SetGamma with it.
func (m *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma((*C.GLFWmonitor)(unsafe.Pointer(m)), C.float(gamma))
}

//GetGammaRamp retrieves the current gamma ramp of the specified monitor.
//
//NOTE: This function does not yet support monitors whose original gamma ramp
//has more or less than gammaRampSize entries.
func (m *Monitor) GetGammaRamp() *GammaRamp {
	var (
		ramp C.GLFWgammaramp
		rampGo GammaRamp
	)
	
	C.glfwGetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), &ramp)
	
	for i := 0; i < GammaRampSize; i++ {
		rampGo.Red[i] = uint16(ramp.red[i])
		rampGo.Green[i] = uint16(ramp.green[i])
		rampGo.Blue[i] = uint16(ramp.blue[i])
	}
	
	return &rampGo
}

//SetGammaRamp sets the current gamma ramp for the specified monitor.
//
//NOTE: This function does not yet support monitors whose original gamma ramp
//has more or less than gammaRampSize entries.
func (m *Monitor) SetGammaRamp(ramp *GammaRamp) {
	var rampC C.GLFWgammaramp
	
	for i := 0; i < GammaRampSize; i++ {
		rampC.red[i] = C.ushort(ramp.Red[i])
		rampC.green[i] = C.ushort(ramp.Green[i])
		rampC.blue[i] = C.ushort(ramp.Blue[i])
	}
	
	C.glfwSetGammaRamp((*C.GLFWmonitor)(unsafe.Pointer(m)), &rampC)
}
