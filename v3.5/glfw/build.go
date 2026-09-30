package glfw

/*
#cgo CFLAGS: -Iglfw/include -D_GNU_SOURCE

// Windows Build Tags
// ----------------
// GLFW Options:
#cgo windows CFLAGS: -D_GLFW_WIN32 -Iglfw/deps/mingw

// Linker Options:
#cgo windows LDFLAGS: -lgdi32

#cgo !gles2,windows LDFLAGS: -lopengl32
#cgo gles2,windows LDFLAGS: -lGLESv2

// Darwin Build Tags
// ----------------
// GLFW Options:
#cgo darwin CFLAGS: -D_GLFW_COCOA -Wno-deprecated-declarations

// Linker Options:
#cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo

#cgo !gles2,darwin LDFLAGS: -framework OpenGL
#cgo gles2,darwin LDFLAGS: -lGLESv2

// Linux Build Tags
// ----------------
// GLFW Options:
// Default (no tags): both X11 and Wayland enabled
// -tags x11: X11 only
// -tags wayland: Wayland only
// -tags x11,wayland: both enabled (explicit)
#cgo linux,!x11,!wayland CFLAGS: -D_GLFW_X11 -D_GLFW_WAYLAND
#cgo linux,x11,!wayland CFLAGS: -D_GLFW_X11
#cgo linux,!x11,wayland CFLAGS: -D_GLFW_WAYLAND
#cgo linux,x11,wayland CFLAGS: -D_GLFW_X11 -D_GLFW_WAYLAND

// Linker Options:
#cgo linux,!gles1,!gles2,!gles3,!vulkan LDFLAGS: -lGL
#cgo linux,gles1 LDFLAGS: -lGLESv1
#cgo linux,gles2 LDFLAGS: -lGLESv2
#cgo linux,gles3 LDFLAGS: -lGLESv3
#cgo linux,vulkan LDFLAGS: -lvulkan
// X11 libraries (included when x11 tag is set or no display tags set)
#cgo linux,!x11,!wayland LDFLAGS: -lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt
#cgo linux,x11 LDFLAGS: -lX11 -lXrandr -lXxf86vm -lXi -lXcursor -lm -lXinerama -ldl -lrt
// Wayland libraries (included when wayland tag is set or no display tags set)
#cgo linux,!x11,!wayland LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon -lm -ldl -lrt
#cgo linux,wayland LDFLAGS: -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon -lm -ldl -lrt

// BSD Build Tags
// ----------------
// GLFW Options:
#cgo freebsd,!wayland netbsd,!wayland openbsd pkg-config: x11 xau xcb xdmcp
#cgo freebsd,wayland netbsd,wayland pkg-config: wayland-client wayland-cursor wayland-egl epoll-shim
#cgo freebsd netbsd openbsd CFLAGS: -D_GLFW_HAS_DLOPEN
#cgo freebsd,!wayland netbsd,!wayland openbsd CFLAGS: -D_GLFW_X11 -D_GLFW_HAS_GLXGETPROCADDRESSARB
#cgo freebsd,wayland netbsd,wayland CFLAGS: -D_GLFW_WAYLAND

// Linker Options:
#cgo freebsd netbsd openbsd LDFLAGS: -lm
*/
import "C"
