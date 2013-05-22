package glfw

//#include <GL/glfw3.h>
//void glfwSetErrorCallbackCB();
import "C"

const (
	//GLFW has not been initialized.
	NotInitialized = C.GLFW_NOT_INITIALIZED
	//No context is current.
	NoCurrentContext = C.GLFW_NO_CURRENT_CONTEXT
	//One of the enum parameters for the function was given an invalid enum.
	InvalidEnum = C.GLFW_INVALID_ENUM
	//One of the parameters for the function was given an invalid value.
	InvalidValue = C.GLFW_INVALID_VALUE
	//A memory allocation failed.
	OutOfMemory = C.GLFW_OUT_OF_MEMORY
	//GLFW could not find support for the requested client API on the system.
	ApiUnavailable = C.GLFW_API_UNAVAILABLE
	//The requested client API version is not available.
	VersionUnavailable = C.GLFW_VERSION_UNAVAILABLE
	//A platform-specific error occurred that does not match any of the more specific categories.
	PlatformError = C.GLFW_PLATFORM_ERROR
	//The clipboard did not contain data in the requested format.
	FormatUnavailable = C.GLFW_FORMAT_UNAVAILABLE
)

type goErrorFunc func(int, string) // Function signature to callback
var fErrorHolder goErrorFunc       // Holds the function for after use

//export goErrorCB
func goErrorCB(err C.int, desc *C.char) {
	fErrorHolder(int(err), C.GoString(desc))
}

//SetErrorCallback sets the error callback, which is called with an error code
//and a human-readable description each time a GLFW error occurs.
//
//This function may be called before Init.
func SetErrorCallback(cbfun goErrorFunc) {
	fErrorHolder = cbfun
	C.glfwSetErrorCallbackCB()
}
