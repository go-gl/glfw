#include <stdint.h>
// #define VK_VERSION_1_0
#include "glfw/deps/vulkan/vulkan.h"
#include "glfw/include/GLFW/glfw3.h"

// Retrieves the address of glfwGetInstanceProcAddress. This workaround is necessary because
// CGO doesn't allow referencing C functions in Go code.
void* getVulkanProcAddr();
