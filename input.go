package glfw

//#include <GL/glfw3.h>
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
	Release = C.GLFW_RELEASE
	Press   = C.GLFW_PRESS
	Repeat  = C.GLFW_REPEAT
)

func (w *Window) GetInputMode(mode int) int {
	return int(C.glfwGetInputMode((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(mode)))
}

func (w *Window) SetInputMode(mode, value int) {
	C.glfwSetInputMode((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(mode), C.int(value))
}

func (w *Window) GetKey(key int) int {
	return int(C.glfwGetKey((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(key)))
}

func (w *Window) GetMouseButton(button int) int {
	return int(C.glfwGetMouseButton((*C.GLFWwindow)(unsafe.Pointer(w)), C.int(button)))
}

func (w *Window) GetCursorPosition() (float64, float64) {
	var xpos, ypos C.double
	C.glfwGetCursorPos((*C.GLFWwindow)(unsafe.Pointer(w)), &xpos, &ypos)
	return float64(xpos), float64(ypos)
}

func (w *Window) SetCursorPosition(xpos, ypos float64) {
	C.glfwSetCursorPos((*C.GLFWwindow)(unsafe.Pointer(w)), C.double(xpos), C.double(ypos))
}

func GetJoystickParameter(joy, param int) int {
	return int(C.glfwGetJoystickParam(C.int(joy), C.int(param)))
}

func GetJoystickAxes(joy int) []float32 {
	var axes [16]C.float
	length := int(C.glfwGetJoystickAxes(C.int(joy), (*C.float)(unsafe.Pointer(&axes[0])), 16))
	a := make([]float32, length)
	for i := 0; i < length; i++ {
		a[i] = float32(axes[i])
	}
	return a
}

func GetJoystickButtons(joy int) []byte {
	var buttons [16]C.uchar
	length := int(C.glfwGetJoystickButtons(C.int(joy), (*C.uchar)(unsafe.Pointer(&buttons[0])), 16))
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = byte(buttons[i])
	}
	return b
}

func GetJoystickName(joy int) string {
	return C.GoString(C.glfwGetJoystickName(C.int(joy)))
}
