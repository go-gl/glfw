//========================================================================
// GLFW 3.1 GLX - www.glfw.org
//------------------------------------------------------------------------
// Copyright (c) 2002-2006 Marcus Geelnard
// Copyright (c) 2006-2010 Camilla Berglund <elmindreda@elmindreda.org>
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgment in the product documentation would
//    be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such, and must not
//    be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source
//    distribution.
//
//========================================================================

#ifndef _glx_context_h_
#define _glx_context_h_

#define GLX_GLXEXT_LEGACY
#include <GL/glx.h>

// This path may need to be changed if you build GLFW using your own setup
// We ship and use our own copy of glxext.h since GLFW uses fairly new
// extensions and not all operating systems come with an up-to-date version
#include "../deps/GL/glxext.h"

// Do we have support for dlopen/dlsym?
#if defined(_GLFW_HAS_DLOPEN)
 #include <dlfcn.h>
#endif

// We support four different ways for getting addresses for GL/GLX
// extension functions: glXGetProcAddress, glXGetProcAddressARB,
// glXGetProcAddressEXT, and dlsym
#if defined(_GLFW_HAS_GLXGETPROCADDRESSARB)
 #define _glfw_glXGetProcAddress(x) glXGetProcAddressARB(x)
#elif defined(_GLFW_HAS_GLXGETPROCADDRESS)
 #define _glfw_glXGetProcAddress(x) glXGetProcAddress(x)
#elif defined(_GLFW_HAS_GLXGETPROCADDRESSEXT)
 #define _glfw_glXGetProcAddress(x) glXGetProcAddressEXT(x)
#elif defined(_GLFW_HAS_DLOPEN)
 #define _glfw_glXGetProcAddress(x) dlsym(_glfw.glx.libGL, x)
 #define _GLFW_DLOPEN_LIBGL
#else
 #error "No OpenGL entry point retrieval mechanism was enabled"
#endif

#define _GLFW_PLATFORM_FBCONFIG                 GLXFBConfig     glx
#define _GLFW_PLATFORM_CONTEXT_STATE            _GLFWcontextGLX glx
#define _GLFW_PLATFORM_LIBRARY_CONTEXT_STATE    _GLFWlibraryGLX glx

#ifndef GLX_MESA_swap_control
typedef int (*PFNGLXSWAPINTERVALMESAPROC)(int);
#endif


// GLX-specific per-context data
//
typedef struct _GLFWcontextGLX
{
    // Rendering context
    GLXContext      context;
    // Visual of selected GLXFBConfig
    XVisualInfo*    visual;

} _GLFWcontextGLX;


// GLX-specific global data
//
typedef struct _GLFWlibraryGLX
{
    int             versionMajor, versionMinor;
    int             eventBase;
    int             errorBase;

    // GLX extensions
    PFNGLXSWAPINTERVALSGIPROC             SwapIntervalSGI;
    PFNGLXSWAPINTERVALEXTPROC             SwapIntervalEXT;
    PFNGLXSWAPINTERVALMESAPROC            SwapIntervalMESA;
    PFNGLXCREATECONTEXTATTRIBSARBPROC     CreateContextAttribsARB;
    GLboolean       SGI_swap_control;
    GLboolean       EXT_swap_control;
    GLboolean       MESA_swap_control;
    GLboolean       ARB_multisample;
    GLboolean       ARB_framebuffer_sRGB;
    GLboolean       ARB_create_context;
    GLboolean       ARB_create_context_profile;
    GLboolean       ARB_create_context_robustness;
    GLboolean       EXT_create_context_es2_profile;
    GLboolean       ARB_context_flush_control;

#if defined(_GLFW_DLOPEN_LIBGL)
    // dlopen handle for libGL.so (for glfwGetProcAddress)
    void*           libGL;
#endif

} _GLFWlibraryGLX;


int _glfwInitContextAPI(void);
void _glfwTerminateContextAPI(void);
int _glfwCreateContext(_GLFWwindow* window,
                       const _GLFWctxconfig* ctxconfig,
                       const _GLFWfbconfig* fbconfig);
void _glfwDestroyContext(_GLFWwindow* window);

#endif // _glx_context_h_
