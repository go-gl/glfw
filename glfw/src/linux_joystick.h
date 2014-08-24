//========================================================================
// GLFW 3.1 Linux - www.glfw.org
//------------------------------------------------------------------------
// Copyright (c) 2014 Jonas Ådahl <jadahl@gmail.com>
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

#ifndef _linux_joystick_h_
#define _linux_joystick_h_

#define _GLFW_PLATFORM_LIBRARY_JOYSTICK_STATE \
    _GLFWjoystickLinux linux_js[GLFW_JOYSTICK_LAST + 1]


//========================================================================
// GLFW platform specific types
//========================================================================

//------------------------------------------------------------------------
// Platform-specific joystick structure
//------------------------------------------------------------------------
typedef struct _GLFWjoystickLinux
{
    int             present;
    int             fd;
    float*          axes;
    int             axisCount;
    unsigned char*  buttons;
    int             buttonCount;
    char*           name;
} _GLFWjoystickLinux;


//========================================================================
// Prototypes for platform specific internal functions
//========================================================================

void _glfwInitJoysticks(void);
void _glfwTerminateJoysticks(void);

#endif // _linux_joystick_h_
