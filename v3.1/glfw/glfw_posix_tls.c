#ifndef GO_GLFW_EXTERNAL
#if defined(_GLFW_COCOA) || defined(_GLFW_X11) || defined(_GLFW_WAYLAND)
	#include "glfw/src/posix_tls.c"
#endif
#endif
