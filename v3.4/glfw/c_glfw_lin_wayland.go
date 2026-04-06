//go:build (linux && !x11 && !wayland) || (linux && wayland)
// +build linux,!x11,!wayland linux,wayland

package glfw

/*
#include "glfw/src/wl_init.c"
#include "glfw/src/wl_monitor.c"
#include "glfw/src/wl_window.c"
*/
import "C"
