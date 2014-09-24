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

// GLFWError holds error code and description.
type GLFWError struct {
	Code ErrorCode
	Desc string
}

// Note: There are many cryptic caveats to proper error handling here.
// See: https://github.com/go-gl/glfw3/pull/86

// Holds the value of the last error.
var lastError = make(chan *GLFWError, 1)

//export goErrorCB
func goErrorCB(code C.int, desc *C.char) {
	flushErrors()
	err := &GLFWError{ErrorCode(code), C.GoString(desc)}
	select {
	case lastError <- err:
	default:
		fmt.Println("GLFW: An uncaught error has occurred:", err)
		fmt.Println("GLFW: Please report this bug in the Go package immediately.")
	}
}

// Error prints the error code and description in a readable format.
func (e *GLFWError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Desc)
}

// Set the glfw callback internally
func init() {
	C.glfwSetErrorCallbackCB()
}

// flushErrors is called by Terminate before it actually calls C.glfwTerminate,
// this ensures that any uncaught errors buffered in lastError are printed
// before the program exits.
func flushErrors() {
	err := fetchError()
	if err != nil {
		fmt.Println("GLFW: An uncaught error has occurred:", err)
		fmt.Println("GLFW: Please report this bug in the Go package immediately.")
	}
}

// fetchError is called by various functions to retrieve the error that might
// have occurred from a generic GLFW operation. It returns nil if no error is
// present.
func fetchError() error {
	select {
	case err := <-lastError:
		return err
	default:
		return nil
	}
}
