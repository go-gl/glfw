//go:build (linux && !wayland) || (linux && x11)
// +build linux,!wayland linux,x11

package glfw

/*
#include "glfw/src/x11_window.c"
#include "glfw/src/x11_init.c"
#include "glfw/src/x11_monitor.c"
*/
import "C"
