package glfw

//#include <GL/glfw3.h>
import "C"

func GetTime() float64 {
	return float64(C.glfwGetTime())
}

func SetTime(time float64) {
	C.glfwSetTime(C.double(time))
}
