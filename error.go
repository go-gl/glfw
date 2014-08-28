package glfw3

//#include "glfw/include/GLFW/glfw3.h"
//void glfwSetErrorCallbackCB();
import "C"

import (
	"fmt"
)

// ErrorCode corresponds to an error code.
type ErrorCode int

// Error codes.
const (
	NotInitialized     ErrorCode = C.GLFW_NOT_INITIALIZED     // GLFW has not been initialized.
	NoCurrentContext   ErrorCode = C.GLFW_NO_CURRENT_CONTEXT  // No context is current.
	InvalidEnum        ErrorCode = C.GLFW_INVALID_ENUM        // One of the enum parameters for the function was given an invalid enum.
	InvalidValue       ErrorCode = C.GLFW_INVALID_VALUE       // One of the parameters for the function was given an invalid value.
	OutOfMemory        ErrorCode = C.GLFW_OUT_OF_MEMORY       // A memory allocation failed.
	APIUnavailable     ErrorCode = C.GLFW_API_UNAVAILABLE     // GLFW could not find support for the requested client API on the system.
	VersionUnavailable ErrorCode = C.GLFW_VERSION_UNAVAILABLE // The requested client API version is not available.
	PlatformError      ErrorCode = C.GLFW_PLATFORM_ERROR      // A platform-specific error occurred that does not match any of the more specific categories.
	FormatUnavailable  ErrorCode = C.GLFW_FORMAT_UNAVAILABLE  // The clipboard did not contain data in the requested format.
)

// GlfwError holds error code and description.
type GLFWError struct {
	Code ErrorCode
	Desc string
}

// Holds the value of the last error
var lastError = make(chan *GLFWError, 1)

//export goErrorCB
func goErrorCB(code C.int, desc *C.char) {
	lastError <- &GLFWError{ErrorCode(code), C.GoString(desc)}
}

// Error prints the error code and description in a readable format.
func (e *GLFWError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Desc)
}

// Set the glfw callback internally
func init() {
	C.glfwSetErrorCallbackCB()
}
