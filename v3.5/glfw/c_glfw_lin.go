//go:build linux
// +build linux

package glfw

/*
#include "glfw/src/glx_context.c"
#include "glfw/src/linux_joystick.c"
#include "glfw/src/posix_module.c"
#include "glfw/src/posix_poll.c"
#include "glfw/src/posix_time.c"
#include "glfw/src/posix_thread.c"
#include "glfw/src/xkb_unicode.c"
#include "glfw/src/egl_context.c"
*/
import "C"
