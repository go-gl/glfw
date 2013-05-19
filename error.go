package glfw

//#include <GL/glfw3.h>
//void glfwSetErrorCallbackCB();
import "C"

const (
	NotInitialized     = C.GLFW_NOT_INITIALIZED
	NoCurrentContext   = C.GLFW_NO_CURRENT_CONTEXT
	InvalidEnum        = C.GLFW_INVALID_ENUM
	InvalidValue       = C.GLFW_INVALID_VALUE
	OutOfMemory        = C.GLFW_OUT_OF_MEMORY
	ApiUnavailable     = C.GLFW_API_UNAVAILABLE
	VersionUnavailable = C.GLFW_VERSION_UNAVAILABLE
	PlatformError      = C.GLFW_PLATFORM_ERROR
	FormatUnavailable  = C.GLFW_FORMAT_UNAVAILABLE
)

type goErrorFunc func(int, string) // Function signature to callback
var fErrorHolder goErrorFunc       // Holds the function for after use

//export goErrorCB
func goErrorCB(err C.int, desc *C.char) {
	fErrorHolder(int(err), C.GoString(desc))
}

func SetErrorCallback(cbfun goErrorFunc) {
	fErrorHolder = cbfun
	C.glfwSetErrorCallbackCB()
}
