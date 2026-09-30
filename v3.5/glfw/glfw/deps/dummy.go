//go:build required
// +build required

// Package dummy prevents go tooling from stripping the c dependencies.
package dummy

import (
	_ "github.com/go-gl/glfw/v3.5/glfw/glfw/deps/glad"
	_ "github.com/go-gl/glfw/v3.5/glfw/glfw/deps/wayland"
)
