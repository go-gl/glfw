package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
//void glfwSetWindowPosCallbackCB(GLFWwindow *window);
//void glfwSetWindowSizeCallbackCB(GLFWwindow *window);
//void glfwSetWindowCloseCallbackCB(GLFWwindow *window);
//void glfwSetWindowRefreshCallbackCB(GLFWwindow *window);
//void glfwSetWindowFocusCallbackCB(GLFWwindow *window);
//void glfwSetWindowIconifyCallbackCB(GLFWwindow *window);
import "C"

import "unsafe"

const (
	Focused   = C.GLFW_FOCUSED
	Iconified = C.GLFW_ICONIFIED
	Visible   = C.GLFW_VISIBLE
	Resizable = C.GLFW_RESIZABLE
	Decorated = C.GLFW_DECORATED

	ClientApi           = C.GLFW_CLIENT_API
	ContextVersionMajor = C.GLFW_CONTEXT_VERSION_MAJOR
	ContextVersionMinor = C.GLFW_CONTEXT_VERSION_MINOR
	ContextRobustness   = C.GLFW_CONTEXT_ROBUSTNESS
	OpenglForwardCompat = C.GLFW_OPENGL_FORWARD_COMPAT
	OpenglDebugContext  = C.GLFW_OPENGL_DEBUG_CONTEXT
	OpenglProfile       = C.GLFW_OPENGL_PROFILE

	OpenglApi   = C.GLFW_OPENGL_API
	OpenglEsApi = C.GLFW_OPENGL_ES_API

	NoRobustness        = C.GLFW_NO_ROBUSTNESS
	NoResetNotification = C.GLFW_NO_RESET_NOTIFICATION
	LoseContextOnReset  = C.GLFW_LOSE_CONTEXT_ON_RESET

	OpenglNoProfile     = C.GLFW_OPENGL_NO_PROFILE
	OpenglCoreProfile   = C.GLFW_OPENGL_CORE_PROFILE
	OpenglCompatProfile = C.GLFW_OPENGL_COMPAT_PROFILE
)

type Window struct {
	data *C.GLFWwindow
}

type (
	goPositionFunc func(*Window, int, int)
	goSizeFunc     func(*Window, int, int)
	goCloseFunc    func(*Window)
	goRefreshFunc  func(*Window)
	goFocusFunc    func(*Window, int)
	goIconifyFunc  func(*Window, int)
)

var (
	fPositionHolder goPositionFunc
	fSizeHolder     goSizeFunc
	fCloseHolder    goCloseFunc
	fRefreshHolder  goRefreshFunc
	fFocusHolder    goFocusFunc
	fIconifyHolder  goIconifyFunc
)

//export goPositionCB
func goPositionCB(window unsafe.Pointer, xpos, ypos C.int) {
	fPositionHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(xpos), int(ypos))
}

//export goSizeCB
func goSizeCB(window unsafe.Pointer, width, height C.int) {
	fSizeHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(width), int(height))
}

//export goCloseCB
func goCloseCB(window unsafe.Pointer) {
	fCloseHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))})
}

//export goRefreshCB
func goRefreshCB(window unsafe.Pointer) {
	fRefreshHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))})
}

//export goFocusCB
func goFocusCB(window unsafe.Pointer, focused C.int) {
	fFocusHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(focused))
}

//export goIconifyCB
func goIconifyCB(window unsafe.Pointer, iconified C.int) {
	fIconifyHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(iconified))
}

func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

func WindowHint(target, hint int) {
	C.glfwWindowHint(C.int(target), C.int(hint))
}

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) *Window {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	return &Window{C.glfwCreateWindow(C.int(width), C.int(height), t, monitor.data, share.data)}
}

func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.data)
}

func (w *Window) ShouldClose() bool {
	r := int(C.glfwWindowShouldClose(w.data))
	if r == C.GL_FALSE {
		return false
	}
	return true
}

func (w *Window) SetShouldClose(value int) {
	C.glfwSetWindowShouldClose(w.data, C.int(value))
}

func (w *Window) SetTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.glfwSetWindowTitle(w.data, t)
}

func (w *Window) GetPosition() (int, int) {
	var (
		xpos C.int
		ypos C.int
	)

	C.glfwGetWindowPos(w.data, &xpos, &ypos)
	return int(xpos), int(ypos)
}

func (w *Window) SetPosition(xpos, ypos int) {
	C.glfwSetWindowPos(w.data, C.int(xpos), C.int(ypos))
}

func (w *Window) GetSize() (int, int) {
	var (
		width  C.int
		height C.int
	)

	C.glfwGetWindowSize(w.data, &width, &height)
	return int(width), int(height)
}

func (w *Window) SetSize(width, height int) {
	C.glfwSetWindowSize(w.data, C.int(width), C.int(height))
}

func (w *Window) Iconify() {
	C.glfwIconifyWindow(w.data)
}

func (w *Window) Restore() {
	C.glfwRestoreWindow(w.data)
}

func (w *Window) Show() {
	C.glfwShowWindow(w.data)
}

func (w *Window) Hide() {
	C.glfwHideWindow(w.data)
}

func (w *Window) GetMonitor() *Monitor {
	return &Monitor{C.glfwGetWindowMonitor(w.data)}
}

func (w *Window) GetAttribute(attrib int) int {
	return int(C.glfwGetWindowAttrib(w.data, C.int(attrib)))
}

func (w *Window) SetUserPointer(pointer unsafe.Pointer) {
	C.glfwSetWindowUserPointer(w.data, pointer)
}

func (w *Window) GetUserPointer() unsafe.Pointer {
	return C.glfwGetWindowUserPointer(w.data)
}

func (w *Window) SetPositionCallback(cbfun goPositionFunc) {
	fPositionHolder = cbfun
	C.glfwSetWindowPosCallbackCB(w.data)
}

func (w *Window) SetSizeCallback(cbfun goSizeFunc) {
	fSizeHolder = cbfun
	C.glfwSetWindowSizeCallbackCB(w.data)
}

func (w *Window) SetCloseCallback(cbfun goCloseFunc) {
	fCloseHolder = cbfun
	C.glfwSetWindowCloseCallbackCB(w.data)
}

func (w *Window) SetRefreshCallback(cbfun goRefreshFunc) {
	fRefreshHolder = cbfun
	C.glfwSetWindowRefreshCallbackCB(w.data)
}

func (w *Window) SetFocusCallback(cbfun goFocusFunc) {
	fFocusHolder = cbfun
	C.glfwSetWindowFocusCallbackCB(w.data)
}

func (w *Window) SetIconifyCallback(cbfun goIconifyFunc) {
	fIconifyHolder = cbfun
	C.glfwSetWindowIconifyCallbackCB(w.data)
}

func PollEvents() {
	C.glfwPollEvents()
}

func WaitEvents() {
	C.glfwWaitEvents()
}
