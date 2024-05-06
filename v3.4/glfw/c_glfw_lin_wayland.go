//go:build linux && wayland
// +build linux,wayland

package glfw

/*
#include "glfw/src/wl_init.c"
#include "glfw/src/wl_monitor.c"
#include "glfw/src/wl_window.c"
*/
import "C"
