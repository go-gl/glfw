package glfw3

//#include <GLFW/glfw3.h>
import "C"

//GetTime returns the value of the GLFW timer. Unless the timer has been set
//using SetTime, the timer measures time elapsed since GLFW was initialized.
//
//The resolution of the timer is system dependent, but is usually on the order
//of a few micro- or nanoseconds. It uses the highest-resolution monotonic time
//source on each supported platform.
func GetTime() (float64, error) {
	r := float64(C.glfwGetTime())
	if r == 0 {
		return 0, <-lastError
	}
	return r, nil
}

//SetTime sets the value of the GLFW timer. It then continues to count up from
//that value.
//
//The resolution of the timer is system dependent, but is usually on the order
//of a few micro- or nanoseconds. It uses the highest-resolution monotonic time
//source on each supported platform.
func SetTime(time float64) {
	C.glfwSetTime(C.double(time))
}
