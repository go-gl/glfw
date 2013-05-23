package glfw

//#include <stdlib.h>
//#include <GL/glfw3.h>
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

type Window C.GLFWwindow

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
	fPositionHolder((*Window)(unsafe.Pointer(window)), int(xpos), int(ypos))
}

//export goSizeCB
func goSizeCB(window unsafe.Pointer, width, height C.int) {
	fSizeHolder((*Window)(unsafe.Pointer(window)), int(width), int(height))
}

//export goCloseCB
func goCloseCB(window unsafe.Pointer) {
	fCloseHolder((*Window)(unsafe.Pointer(window)))
}

//export goRefreshCB
func goRefreshCB(window unsafe.Pointer) {
	fRefreshHolder((*Window)(unsafe.Pointer(window)))
}

//export goFocusCB
func goFocusCB(window unsafe.Pointer, focused C.int) {
	fFocusHolder((*Window)(unsafe.Pointer(window)), int(focused))
}

//export goIconifyCB
func goIconifyCB(window unsafe.Pointer, iconified C.int) {
	fIconifyHolder((*Window)(unsafe.Pointer(window)), int(iconified))
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
	return (*Window)(unsafe.Pointer(C.glfwCreateWindow(C.int(width), C.int(height), t, (*C.GLFWmonitor)(unsafe.Pointer(monitor)), (*C.GLFWwindow)(unsafe.Pointer(share)))))
}

func (w *Window) Destroy() {
	C.glfwDestroyWindow((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) ShouldClose() bool {
	r := int(C.glfwWindowShouldClose((*C.GLFWwindow)(unsafe.Pointer(w))))
	if r == C.GL_FALSE {
		return false
	}
	return true
}

func (w *Window) SetShouldClose(value int) {
	C.glfwSetWindowShouldClose((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(value))
}

func (w *Window) SetTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.glfwSetWindowTitle((*C.GLFWwindow)(unsafe.Pointer(w)), t)
}

func (w *Window) GetPosition() (int, int) {
	var (
		xpos C.int
		ypos C.int
	)
	C.glfwGetWindowPos((*C.GLFWwindow)(unsafe.Pointer(w)), &xpos, &ypos)
	return int(xpos), int(ypos)
}

func (w *Window) SetPosition(xpos, ypos int) {
	C.glfwSetWindowPos((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(xpos), C.int(ypos))
}

func (w *Window) GetSize() (int, int) {
	var (
		width  C.int
		height C.int
	)
	C.glfwGetWindowSize((*C.GLFWwindow)(unsafe.Pointer(w)), &width, &height)
	return int(width), int(height)
}

func (w *Window) SetSize(width, height int) {
	C.glfwSetWindowSize((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(width), C.int(height))
}

func (w *Window) Iconify() {
	C.glfwIconifyWindow((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) Restore() {
	C.glfwRestoreWindow((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) Show() {
	C.glfwShowWindow((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) Hide() {
	C.glfwHideWindow((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) GetMonitor() *Monitor {
	return (*Monitor)(unsafe.Pointer(C.glfwGetWindowMonitor((*C.GLFWwindow)(unsafe.Pointer(w)))))
}

func (w *Window) GetParameter(param int) int {
	return int(C.glfwGetWindowParam((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(param)))
}

func (w *Window) SetUserPointer(pointer unsafe.Pointer) {
	C.glfwSetWindowUserPointer((*C.GLFWwindow)(unsafe.Pointer(w)), pointer)
}

func (w *Window) GetUserPointer() unsafe.Pointer {
	return C.glfwGetWindowUserPointer((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetPositionCallback(cbfun goPositionFunc) {
	fPositionHolder = cbfun
	C.glfwSetWindowPosCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetSizeCallback(cbfun goSizeFunc) {
	fSizeHolder = cbfun
	C.glfwSetWindowSizeCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetCloseCallback(cbfun goCloseFunc) {
	fCloseHolder = cbfun
	C.glfwSetWindowCloseCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetRefreshCallback(cbfun goRefreshFunc) {
	fRefreshHolder = cbfun
	C.glfwSetWindowRefreshCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetFocusCallback(cbfun goFocusFunc) {
	fFocusHolder = cbfun
	C.glfwSetWindowFocusCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func (w *Window) SetIconifyCallback(cbfun goIconifyFunc) {
	fIconifyHolder = cbfun
	C.glfwSetWindowIconifyCallbackCB((*C.GLFWwindow)(unsafe.Pointer(w)))
}

func PollEvents() {
	C.glfwPollEvents()
}

func WaitEvents() {
	C.glfwWaitEvents()
}
