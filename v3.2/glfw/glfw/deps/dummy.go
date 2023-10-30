//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/go-gl/glfw/v3.2/glfw/glfw/deps/KHR"
	_ "github.com/go-gl/glfw/v3.2/glfw/glfw/deps/glad"
	_ "github.com/go-gl/glfw/v3.2/glfw/glfw/deps/mingw"
	_ "github.com/go-gl/glfw/v3.2/glfw/glfw/deps/vulkan"
)
