package glfw

//#include <stdlib.h>
//#define GLFW_INCLUDE_NONE
//#include "glfw/include/GLFW/glfw3.h"
import "C"

// Platform is a platform GLFW supports.
type Platform int

const (
	Win32Platform   Platform = C.GLFW_PLATFORM_WIN32
	CocoaPlatform   Platform = C.GLFW_PLATFORM_COCOA
	WaylandPlatform Platform = C.GLFW_PLATFORM_WAYLAND
	X11Platform     Platform = C.GLFW_PLATFORM_X11
	NullPlatform    Platform = C.GLFW_PLATFORM_NULL
)

// GetPlatform returns the currently selected platform.
func GetPlatform() Platform {
	return Platform(C.glfwGetPlatform())
}

// Supported returns whether the library includes support for the specified platform.
func (p Platform) Suported() bool {
	return C.glfwPlatformSupported(C.int(p)) == 1
}
