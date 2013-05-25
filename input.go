package glfw

//#include <GLFW/glfw3.h>
//void glfwSetKeyCallbackCB(GLFWwindow *window);
//void glfwSetCharCallbackCB(GLFWwindow *window);
//void glfwSetMouseCallbackCB(GLFWwindow *window);
//void glfwSetPosCallbackCB(GLFWwindow *window);
//void glfwSetEnterCallbackCB(GLFWwindow *window);
//void glfwSetScrollCallbackCB(GLFWwindow *window);
//float GetAxisAtIndex(float *axis, int i);
//unsigned char GetButtonsAtIndex(unsigned char *buttons, int i);
import "C"

import "unsafe"

const (
	Joystick1    = C.GLFW_JOYSTICK_1
	Joystick2    = C.GLFW_JOYSTICK_2
	Joystick3    = C.GLFW_JOYSTICK_3
	Joystick4    = C.GLFW_JOYSTICK_4
	Joystick5    = C.GLFW_JOYSTICK_5
	Joystick6    = C.GLFW_JOYSTICK_6
	Joystick7    = C.GLFW_JOYSTICK_7
	Joystick8    = C.GLFW_JOYSTICK_8
	Joystick9    = C.GLFW_JOYSTICK_9
	Joystick10   = C.GLFW_JOYSTICK_10
	Joystick11   = C.GLFW_JOYSTICK_11
	Joystick12   = C.GLFW_JOYSTICK_12
	Joystick13   = C.GLFW_JOYSTICK_13
	Joystick14   = C.GLFW_JOYSTICK_14
	Joystick15   = C.GLFW_JOYSTICK_15
	Joystick16   = C.GLFW_JOYSTICK_16
	JoystickLast = C.GLFW_JOYSTICK_LAST
)

//These key codes are inspired by the USB HID Usage Tables v1.12 (p. 53-60),
//but re-arranged to map to 7-bit ASCII for printable keys (function keys are
//put in the 256+ range).
const (
	KeySpace        = C.GLFW_KEY_SPACE
	KeyApostrophe   = C.GLFW_KEY_APOSTROPHE
	KeyComma        = C.GLFW_KEY_COMMA
	KeyMinus        = C.GLFW_KEY_MINUS
	KeyPeriod       = C.GLFW_KEY_PERIOD
	KeySlash        = C.GLFW_KEY_SLASH
	Key0            = C.GLFW_KEY_0
	Key1            = C.GLFW_KEY_1
	Key2            = C.GLFW_KEY_2
	Key3            = C.GLFW_KEY_3
	Key4            = C.GLFW_KEY_4
	Key5            = C.GLFW_KEY_5
	Key6            = C.GLFW_KEY_6
	Key7            = C.GLFW_KEY_7
	Key8            = C.GLFW_KEY_8
	Key9            = C.GLFW_KEY_9
	KeySemicolon    = C.GLFW_KEY_SEMICOLON
	KeyEqual        = C.GLFW_KEY_EQUAL
	KeyA            = C.GLFW_KEY_A
	KeyB            = C.GLFW_KEY_B
	KeyC            = C.GLFW_KEY_C
	KeyD            = C.GLFW_KEY_D
	KeyE            = C.GLFW_KEY_E
	KeyF            = C.GLFW_KEY_F
	KeyG            = C.GLFW_KEY_G
	KeyH            = C.GLFW_KEY_H
	KeyI            = C.GLFW_KEY_I
	KeyJ            = C.GLFW_KEY_J
	KeyK            = C.GLFW_KEY_K
	KeyL            = C.GLFW_KEY_L
	KeyM            = C.GLFW_KEY_M
	KeyN            = C.GLFW_KEY_N
	KeyO            = C.GLFW_KEY_O
	KeyP            = C.GLFW_KEY_P
	KeyQ            = C.GLFW_KEY_Q
	KeyR            = C.GLFW_KEY_R
	KeyS            = C.GLFW_KEY_S
	KeyT            = C.GLFW_KEY_T
	KeyU            = C.GLFW_KEY_U
	KeyV            = C.GLFW_KEY_V
	KeyW            = C.GLFW_KEY_W
	KeyX            = C.GLFW_KEY_X
	KeyY            = C.GLFW_KEY_Y
	KeyZ            = C.GLFW_KEY_Z
	KeyLeftBracket  = C.GLFW_KEY_LEFT_BRACKET
	KeyBackslash    = C.GLFW_KEY_BACKSLASH
	KeyBracket      = C.GLFW_KEY_RIGHT_BRACKET
	KeyGraveAccent  = C.GLFW_KEY_GRAVE_ACCENT
	KeyWorld1       = C.GLFW_KEY_WORLD_1
	KeyWorld2       = C.GLFW_KEY_WORLD_2
	KeyEscape       = C.GLFW_KEY_ESCAPE
	KeyEnter        = C.GLFW_KEY_ENTER
	KeyTab          = C.GLFW_KEY_TAB
	KeyBackspace    = C.GLFW_KEY_BACKSPACE
	KeyInsert       = C.GLFW_KEY_INSERT
	KeyDelete       = C.GLFW_KEY_DELETE
	KeyRight        = C.GLFW_KEY_RIGHT
	KeyLeft         = C.GLFW_KEY_LEFT
	KeyDown         = C.GLFW_KEY_DOWN
	KeyUp           = C.GLFW_KEY_UP
	KeyPageUp       = C.GLFW_KEY_PAGE_UP
	KeyPageDown     = C.GLFW_KEY_PAGE_DOWN
	KeyHome         = C.GLFW_KEY_HOME
	KeyEnd          = C.GLFW_KEY_END
	KeyCapsLock     = C.GLFW_KEY_CAPS_LOCK
	KeyScrollLock   = C.GLFW_KEY_SCROLL_LOCK
	KeyNumLock      = C.GLFW_KEY_NUM_LOCK
	KeyPrintScreen  = C.GLFW_KEY_PRINT_SCREEN
	KeyPause        = C.GLFW_KEY_PAUSE
	KeyF1           = C.GLFW_KEY_F1
	KeyF2           = C.GLFW_KEY_F2
	KeyF3           = C.GLFW_KEY_F3
	KeyF4           = C.GLFW_KEY_F4
	KeyF5           = C.GLFW_KEY_F5
	KeyF6           = C.GLFW_KEY_F6
	KeyF7           = C.GLFW_KEY_F7
	KeyF8           = C.GLFW_KEY_F8
	KeyF9           = C.GLFW_KEY_F9
	KeyF10          = C.GLFW_KEY_F10
	KeyF11          = C.GLFW_KEY_F11
	KeyF12          = C.GLFW_KEY_F12
	KeyF13          = C.GLFW_KEY_F13
	KeyF14          = C.GLFW_KEY_F14
	KeyF15          = C.GLFW_KEY_F15
	KeyF16          = C.GLFW_KEY_F16
	KeyF17          = C.GLFW_KEY_F17
	KeyF18          = C.GLFW_KEY_F18
	KeyF19          = C.GLFW_KEY_F19
	KeyF20          = C.GLFW_KEY_F20
	KeyF21          = C.GLFW_KEY_F21
	KeyF22          = C.GLFW_KEY_F22
	KeyF23          = C.GLFW_KEY_F23
	KeyF24          = C.GLFW_KEY_F24
	KeyF25          = C.GLFW_KEY_F25
	KeyKp0          = C.GLFW_KEY_KP_0
	KeyKp1          = C.GLFW_KEY_KP_1
	KeyKp2          = C.GLFW_KEY_KP_2
	KeyKp3          = C.GLFW_KEY_KP_3
	KeyKp4          = C.GLFW_KEY_KP_4
	KeyKp5          = C.GLFW_KEY_KP_5
	KeyKp6          = C.GLFW_KEY_KP_6
	KeyKp7          = C.GLFW_KEY_KP_7
	KeyKp8          = C.GLFW_KEY_KP_8
	KeyKp9          = C.GLFW_KEY_KP_9
	KeyKpDecimal    = C.GLFW_KEY_KP_DECIMAL
	KeyKpDivide     = C.GLFW_KEY_KP_DIVIDE
	KeyKpMultiply   = C.GLFW_KEY_KP_MULTIPLY
	KeyKpSubtract   = C.GLFW_KEY_KP_SUBTRACT
	KeyKpAdd        = C.GLFW_KEY_KP_ADD
	KeyKpEnter      = C.GLFW_KEY_KP_ENTER
	KeyKpEqual      = C.GLFW_KEY_KP_EQUAL
	KeyLeftShift    = C.GLFW_KEY_LEFT_SHIFT
	KeyLeftControl  = C.GLFW_KEY_LEFT_CONTROL
	KeyLeftAlt      = C.GLFW_KEY_LEFT_ALT
	KeyLeftSuper    = C.GLFW_KEY_LEFT_SUPER
	KeyRightShift   = C.GLFW_KEY_RIGHT_SHIFT
	KeyRightControl = C.GLFW_KEY_RIGHT_CONTROL
	KeyRightAlt     = C.GLFW_KEY_RIGHT_ALT
	KeyRightSuper   = C.GLFW_KEY_RIGHT_SUPER
	KeyMenu         = C.GLFW_KEY_MENU
	KeyLast         = C.GLFW_KEY_LAST
)

const (
	MouseButton1      = C.GLFW_MOUSE_BUTTON_1
	MouseButton2      = C.GLFW_MOUSE_BUTTON_2
	MouseButton3      = C.GLFW_MOUSE_BUTTON_3
	MouseButton4      = C.GLFW_MOUSE_BUTTON_4
	MouseButton5      = C.GLFW_MOUSE_BUTTON_5
	MouseButton6      = C.GLFW_MOUSE_BUTTON_6
	MouseButton7      = C.GLFW_MOUSE_BUTTON_7
	MouseButton8      = C.GLFW_MOUSE_BUTTON_8
	MouseButtonLast   = C.GLFW_MOUSE_BUTTON_LAST
	MouseButtonLeft   = C.GLFW_MOUSE_BUTTON_LEFT
	MouseButtonRight  = C.GLFW_MOUSE_BUTTON_RIGHT
	MouseButtonMiddle = C.GLFW_MOUSE_BUTTON_MIDDLE
)

const (
	//The key or button was released.
	Release = C.GLFW_RELEASE
	//The key or button was pressed.
	Press = C.GLFW_PRESS
	//The key was held down until it repeated.
	Repeat = C.GLFW_REPEAT
)

//Input modes
const (
	//See Cursor mode values
	Cursor = C.GLFW_CURSOR
	//Value can be either 1 or 0
	StickyKeys = C.GLFW_STICKY_KEYS
	//Value can be either 1 or 0
	StickyMouseButtons = C.GLFW_STICKY_MOUSE_BUTTONS
)

//Cursor mode values
const (
	CursorNormal   = C.GLFW_CURSOR_NORMAL
	CursorHidden   = C.GLFW_CURSOR_HIDDEN
	CursorDisabled = C.GLFW_CURSOR_DISABLED
)

type (
	goMouseFunc  func(*Window, int, int, int)
	goPosFunc    func(*Window, float64, float64)
	goEnterFunc  func(*Window, int)
	goScrollFunc func(*Window, float64, float64)
	goKeyFunc    func(*Window, int, int, int)
	goCharFunc   func(*Window, uint)
)

var (
	fMouseHolder  goMouseFunc
	fPosHolder    goPosFunc
	fEnterHolder  goEnterFunc
	fScrollHolder goScrollFunc
	fKeyHolder    goKeyFunc
	fCharHolder   goCharFunc
)

//export goMouseCB
func goMouseCB(window unsafe.Pointer, button, action, mods C.int) {
	fMouseHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(button), int(action), int(mods))
}

//export goPosCB
func goPosCB(window unsafe.Pointer, xpos, ypos C.double) {
	fPosHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, float64(xpos), float64(ypos))
}

//export goEnterCB
func goEnterCB(window unsafe.Pointer, entered C.int) {
	fEnterHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(entered))
}

//export goScrollCB
func goScrollCB(window unsafe.Pointer, xpos, ypos C.double) {
	fScrollHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, float64(xpos), float64(ypos))
}

//export goKeyCB
func goKeyCB(window unsafe.Pointer, key, action, mods C.int) {
	fKeyHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(key), int(action), int(mods))
}

//export goCharCB
func goCharCB(window unsafe.Pointer, character C.uint) {
	fCharHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, uint(character))
}

//GetInputMode returns the value of an input option.
func (w *Window) GetInputMode(mode int) int {
	return int(C.glfwGetInputMode(w.data, C.int(mode)))
}

//Sets an input option.
func (w *Window) SetInputMode(mode, value int) {
	C.glfwSetInputMode(w.data, C.int(mode), C.int(value))
}

//GetKey returns the last reported state of a keyboard key. The returned state
//is one of Press or Release. The higher-level state Repeat is only reported to
//the key callback.
//
//If the StickyKeys input mode is enabled, this function returns Press the first
//time you call this function after a key has been pressed, even if the key has
//already been released.
//
//The key functions deal with physical keys, with key tokens named after their
//use on the standard US keyboard layout. If you want to input text, use the
//Unicode character callback instead.
func (w *Window) GetKey(key int) int {
	return int(C.glfwGetKey(w.data, C.int(key)))
}

//GetMouseButton returns the last state reported for the specified mouse button.
//
//If the StickyMouseButtons input mode is enabled, this function returns Press
//the first time you call this function after a mouse button has been pressed,
//even if the mouse button has already been released.
func (w *Window) GetMouseButton(button int) int {
	return int(C.glfwGetMouseButton(w.data, C.int(button)))
}

//GetCursorPosition returns the last reported position of the cursor.
//
//If the cursor is disabled (with CursorDisabled) then the cursor position is
//unbounded and limited only by the minimum and maximum values of a double.
//
//The coordinate can be converted to their integer equivalents with the floor
//function. Casting directly to an integer type works for positive coordinates,
//but fails for negative ones.
func (w *Window) GetCursorPosition() (float64, float64) {
	var xpos, ypos C.double
	C.glfwGetCursorPos(w.data, &xpos, &ypos)
	return float64(xpos), float64(ypos)
}

//SetCursorPosition sets the position of the cursor. The specified window mus
//be focused. If the window does not have focus when this function is called,
//it fails silently.
//
//If the cursor is disabled (with CursorDisabled) then the cursor position is
//unbounded and limited only by the minimum and maximum values of a double.
func (w *Window) SetCursorPosition(xpos, ypos float64) {
	C.glfwSetCursorPos(w.data, C.double(xpos), C.double(ypos))
}

//SetKeyCallback sets the key callback which is called when a key is pressed,
//repeated or released.
//
//The key functions deal with physical keys, with layout independent key tokens
//named after their values in the standard US keyboard layout. If you want to
//input text, use the SetCharCallback instead.
//
//When a window loses focus, it will generate synthetic key release events for
//all pressed keys. You can tell these events from user-generated events by the
//fact that the synthetic ones are generated after the window has lost focus,
//i.e. Focused will be false and the focus callback will have already been
//called.
//
//Function signature for this callback is: func(*Window, int, int)
func (w *Window) SetKeyCallback(cbfun goKeyFunc) {
	fKeyHolder = cbfun
	C.glfwSetKeyCallbackCB(w.data)
}

//SetCharacterCallback sets the character callback which is called when a
//Unicode character is input.
//
//The character callback is intended for text input. If you want to know whether
//a specific key was pressed or released, use the key callback instead.
//
//Function signature for this callback is: func(*Window, uint)
func (w *Window) SetCharacterCallback(cbfun goCharFunc) {
	fCharHolder = cbfun
	C.glfwSetCharCallbackCB(w.data)
}

//SetMouseButtonCallback sets the mouse button callback which is called when a
//mouse button is pressed or released.
//
//When a window loses focus, it will generate synthetic mouse button release
//events for all pressed mouse buttons. You can tell these events from
//user-generated events by the fact that the synthetic ones are generated after
//the window has lost focus, i.e. Focused will be false and the focus
//callback will have already been called.
//
//Function signature for this callback is: func(*Window, int, int)
func (w *Window) SetMouseButtonCallback(cbfun goMouseFunc) {
	fMouseHolder = cbfun
	C.glfwSetMouseCallbackCB(w.data)
}

//SetCursorPositionCallback sets the cursor position callback which is called
//when the cursor is moved. The callback is provided
//with the position relative to the upper-left corner of the client area of the
//window.
//
//Function signature for this callback is: func(*Window, float64, float64)
func (w *Window) SetCursorPositionCallback(cbfun goPosFunc) {
	fPosHolder = cbfun
	C.glfwSetPosCallbackCB(w.data)
}

//SetCursorEnterCallback the cursor boundary crossing callback which is called
//when the cursor enters or leaves the client area of the window.
//
//Function signature for this callback is: func(*Window, int)
func (w *Window) SetCursorEnterCallback(cbfun goEnterFunc) {
	fEnterHolder = cbfun
	C.glfwSetEnterCallbackCB(w.data)
}

//SetScrollCallback sets the scroll callback which is called when a scrolling
//device is used, such as a mouse wheel or scrolling area of a touchpad.
//
//Function signature for this callback is: func(*Window, float64, float64)
func (w *Window) SetScrollCallback(cbfun goScrollFunc) {
	fScrollHolder = cbfun
	C.glfwSetScrollCallbackCB(w.data)
}

//GetJoystickPresent returns whether the specified joystick is present.
func JoystickPresent(joy int) int {
	return int(C.glfwJoystickPresent(C.int(joy)))
}

//GetJoystickAxes returns an array of axis values.
func GetJoystickAxes(joy int) []float32 {
	var length int
	axis := C.glfwGetJoystickAxes(C.int(joy), (*C.int)(unsafe.Pointer(&length)))
	a := make([]float32, length)
	for i := 0; i < length; i++ {
		a[i] = float32(C.GetAxisAtIndex(axis, C.int(i)))
	}
	return a
}

func GetJoystickButtons(joy int) []byte {
	var length int
	buttons := C.glfwGetJoystickButtons(C.int(joy), (*C.int)(unsafe.Pointer(&length)))
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(C.GetButtonsAtIndex(buttons, C.int(i)))
	}
	return b
}

func GetJoystickName(joy int) string {
	return C.GoString(C.glfwGetJoystickName(C.int(joy)))
}
