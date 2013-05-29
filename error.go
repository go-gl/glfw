package glfw

//#include <GLFW/glfw3.h>
//void glfwSetErrorCallbackCB();
import "C"

const (
	NotInitialized     = C.GLFW_NOT_INITIALIZED     //GLFW has not been initialized.
	NoCurrentContext   = C.GLFW_NO_CURRENT_CONTEXT  //No context is current.
	InvalidEnum        = C.GLFW_INVALID_ENUM        //One of the enum parameters for the function was given an invalid enum.
	InvalidValue       = C.GLFW_INVALID_VALUE       //One of the parameters for the function was given an invalid value.
	OutOfMemory        = C.GLFW_OUT_OF_MEMORY       //A memory allocation failed.
	ApiUnavailable     = C.GLFW_API_UNAVAILABLE     //GLFW could not find support for the requested client API on the system.
	VersionUnavailable = C.GLFW_VERSION_UNAVAILABLE //The requested client API version is not available.
	PlatformError      = C.GLFW_PLATFORM_ERROR      //A platform-specific error occurred that does not match any of the more specific categories.
	FormatUnavailable  = C.GLFW_FORMAT_UNAVAILABLE  //The clipboard did not contain data in the requested format.
)

type goErrorFunc func(int, string)

var fErrorHolder goErrorFunc

//export goErrorCB
func goErrorCB(err C.int, desc *C.char) {
	fErrorHolder(int(err), C.GoString(desc))
}

//SetErrorCallback sets the error callback, which is called with an error code
//and a human-readable description each time a GLFW error occurs.
//
//This function may be called before Init.
//
//Function signature for this callback is: func(int, string)
func SetErrorCallback(cbfun goErrorFunc) {
	fErrorHolder = cbfun
	C.glfwSetErrorCallbackCB()
}
