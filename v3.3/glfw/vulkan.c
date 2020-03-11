#include "vulkan.h"

void* getVulkanProcAddr() {
	return glfwGetInstanceProcAddress;
}
