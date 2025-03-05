//go:build !((linux && wayland) || (freebsd && wayland) || (netbsd && wayland) || (openbsd && wayland))
// -build linux,wayland freebsd,wayland netbsd,wayland openbsd,wayland

package glfw

const WAYLAND = false