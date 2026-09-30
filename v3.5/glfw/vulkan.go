package glfw

/*
#include "glfw/src/internal.h"

GLFWAPI VkResult glfwCreateWindowSurface(VkInstance instance, GLFWwindow* window, const VkAllocationCallbacks* allocator, VkSurfaceKHR* surface);
GLFWAPI GLFWvkproc glfwGetInstanceProcAddress(VkInstance instance, const char* procname);
GLFWAPI void glfwInitVulkanLoader(PFN_vkGetInstanceProcAddr loader);
GLFWAPI int glfwGetPhysicalDevicePresentationSupport(VkInstance instance, VkPhysicalDevice device, uint32_t queuefamily);

// Helper function for doing raw pointer arithmetic
static inline const char* getArrayIndex(const char** array, unsigned int index) {
	return array[index];
}

void* getVulkanProcAddr() {
	return glfwGetInstanceProcAddress;
}

static inline void glfwInitVulkanLoaderBridge(void* loader) {
	glfwInitVulkanLoader((PFN_vkGetInstanceProcAddr) loader);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// VulkanSupported reports whether the Vulkan loader has been found. This check is performed by Init.
//
// The availability of a Vulkan loader does not by itself guarantee that window surface creation or
// even device creation is possible. Call GetRequiredInstanceExtensions to check whether the
// extensions necessary for Vulkan surface creation are available and GetPhysicalDevicePresentationSupport
// to check whether a queue family of a physical device supports image presentation.
func VulkanSupported() bool {
	return glfwbool(C.glfwVulkanSupported())
}

// GetVulkanGetInstanceProcAddress returns the function pointer used to find Vulkan core or
// extension functions. The return value of this function can be passed to the Vulkan library.
//
// Note that this function does not work the same way as the glfwGetInstanceProcAddress.
func GetVulkanGetInstanceProcAddress() unsafe.Pointer {
	return C.getVulkanProcAddr()
}

// InitVulkanLoader sets the Vulkan loader function for the next call to Init.
//
// Pass nil to use GLFW's default dynamic loader behavior.
//
// This function must only be called from the main thread.
func InitVulkanLoader(loader unsafe.Pointer) {
	C.glfwInitVulkanLoaderBridge(loader)
	panicError()
}

// GetPhysicalDevicePresentationSupport reports whether a queue family of a
// physical device supports presentation to the active GLFW platform.
func GetPhysicalDevicePresentationSupport(instance, device interface{}, queueFamily uint32) (supported bool, _ error) {
	if instance == nil {
		return false, errors.New("vulkan: instance is nil")
	}
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() != reflect.Ptr {
		return false, fmt.Errorf("vulkan: instance is not a VkInstance (expected kind Ptr, got %s)", instanceValue.Kind())
	}

	if device == nil {
		return false, errors.New("vulkan: device is nil")
	}
	deviceValue := reflect.ValueOf(device)
	if deviceValue.Kind() != reflect.Ptr {
		return false, fmt.Errorf("vulkan: device is not a VkPhysicalDevice (expected kind Ptr, got %s)", deviceValue.Kind())
	}

	supported = glfwbool(C.glfwGetPhysicalDevicePresentationSupport(
		(C.VkInstance)(unsafe.Pointer(instanceValue.Pointer())),
		(C.VkPhysicalDevice)(unsafe.Pointer(deviceValue.Pointer())),
		C.uint32_t(queueFamily),
	))
	if err := acceptError(APIUnavailable); err != nil {
		return false, err
	}
	panicError()
	return supported, nil
}

// GetRequiredInstanceExtensions returns a slice of Vulkan instance extension names required
// by GLFW for creating Vulkan surfaces for GLFW windows. If successful, the list will always
// contain VK_KHR_surface, so if you don't require any additional extensions you can pass this list
// directly to the VkInstanceCreateInfo struct.
//
// If Vulkan is not available on the machine, this function returns nil. Call
// VulkanSupported to check whether Vulkan is available.
//
// If Vulkan is available but no set of extensions allowing window surface creation was found, this
// function returns nil. You may still use Vulkan for off-screen rendering and compute work.
func (window *Window) GetRequiredInstanceExtensions() []string {
	var count C.uint32_t
	strarr := C.glfwGetRequiredInstanceExtensions(&count)
	if count == 0 {
		return nil
	}

	extensions := make([]string, count)
	for i := uint(0); i < uint(count); i++ {
		extensions[i] = C.GoString(C.getArrayIndex(strarr, C.uint(i)))
	}
	return extensions
}

// CreateWindowSurface creates a Vulkan surface for this window.
func (window *Window) CreateWindowSurface(instance interface{}, allocCallbacks unsafe.Pointer) (surface uintptr, _ error) {
	if instance == nil {
		return 0, errors.New("vulkan: instance is nil")
	}
	val := reflect.ValueOf(instance)
	if val.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("vulkan: instance is not a VkInstance (expected kind Ptr, got %s)", val.Kind())
	}
	var vulkanSurface C.VkSurfaceKHR
	ret := C.glfwCreateWindowSurface(
		(C.VkInstance)(unsafe.Pointer(reflect.ValueOf(instance).Pointer())), window.data,
		(*C.VkAllocationCallbacks)(allocCallbacks), (*C.VkSurfaceKHR)(unsafe.Pointer(&vulkanSurface)))
	if ret != C.VK_SUCCESS {
		return 0, fmt.Errorf("vulkan: error creating window surface: %d", ret)
	}
	return uintptr(unsafe.Pointer(&vulkanSurface)), nil
}
