package glfw3

//#include "glfw/include/GLFW/glfw3.h"
//void glfwSetErrorCallbackCB();
import "C"

import (
	"fmt"
)

// prefixedError is a error that returns a error string in the format of:
//
//  prefix: desc
//
type prefixedError struct {
	prefix, desc string
}

// Error implements the error interface.
func (e prefixedError) Error() string {
	return fmt.Sprintf("%s: %s", e.prefix, e.desc)
}

var (
	// ErrNotInitialized is an error returned when GLFW has not yet been
	// initialized.
	ErrNotInitialized = prefixedError{prefix: "ErrNotInitialized"}

	// ErrNoCurrentContext is an error returned when no OpenGL context has
	// been made current in the OS thread.
	ErrNoCurrentContext = prefixedError{prefix: "ErrNoCurrentContext"}

	// ErrInvalidEnum is an error returned when you've passed an invalid enum
	// to a function as a parameter.
	ErrInvalidEnum = prefixedError{prefix: "ErrInvalidEnum"}

	// ErrInvalidValue is an error returned when you've passed an invalid value
	// to a function as a parameter.
	ErrInvalidValue = prefixedError{prefix: "ErrInvalidValue"}

	// ErrOutOfMemory is an error returned when GLFW has ran out of memory and
	// allocation failed.
	ErrOutOfMemory = prefixedError{prefix: "ErrOutOfMemory"}

	// ErrAPIUnavailable is an error returned when GLFW could not find support
	// for the requested client API on the system.
	ErrAPIUnavailable = prefixedError{prefix: "ErrAPIUnavailable"}

	// ErrVersionUnavailable is an error returned when the requested client API
	// version is not available.
	ErrVersionUnavailable = prefixedError{prefix: "ErrVersionUnavailable"}

	// ErrPlatformError is an error returned when a platform-specific error
	// occurred that does not match any of the more specific categories.
	ErrPlatformError = prefixedError{prefix: "ErrPlatformError"}

	// ErrFormatUnavailable is an error returned when the clipboard did not
	// contain data in the requested format.
	ErrFormatUnavailable = prefixedError{prefix: "ErrFormatUnavailable"}
)

// newError finds the error associated with the code and returns an appropriate
// Go error value.
func newError(code C.int, desc string) error {
	var p prefixedError
	switch code {
	case C.GLFW_NOT_INITIALIZED:
		p = ErrNotInitialized
	case C.GLFW_NO_CURRENT_CONTEXT:
		p = ErrNoCurrentContext
	case C.GLFW_INVALID_ENUM:
		p = ErrInvalidEnum
	case C.GLFW_INVALID_VALUE:
		p = ErrInvalidValue
	case C.GLFW_OUT_OF_MEMORY:
		p = ErrOutOfMemory
	case C.GLFW_API_UNAVAILABLE:
		p = ErrAPIUnavailable
	case C.GLFW_VERSION_UNAVAILABLE:
		p = ErrVersionUnavailable
	case C.GLFW_PLATFORM_ERROR:
		p = ErrPlatformError
	case C.GLFW_FORMAT_UNAVAILABLE:
		p = ErrFormatUnavailable
	default:
		panic(fmt.Sprintf("unknown error code (0x%X)", code))
	}
	p.desc = desc
	return p
}

// Note: There are many cryptic caveats to proper error handling here.
// See: https://github.com/go-gl/glfw3/pull/86

// Holds the value of the last error.
var lastError = make(chan error, 1)

//export goErrorCB
func goErrorCB(code C.int, desc *C.char) {
	flushErrors()
	err := newError(code, C.GoString(desc))
	select {
	case lastError <- err:
	default:
		fmt.Println("GLFW: An uncaught error has occurred:", err)
		fmt.Println("GLFW: Please report this bug in the Go package immediately.")
	}
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
