package glfw

//#cgo LDFLAGS: -lglfw
//#include <GL/glfw3.h>
import "C"

const (
	VersionMajor    = C.GLFW_VERSION_MAJOR
	VersionMinor    = C.GLFW_VERSION_MINOR
	VersionRevision = C.GLFW_VERSION_REVISION
)

func Init() bool {
	r := C.glfwInit()
	if r == C.GL_TRUE {
		return true
	}
	return false
}

func Terminate() {
	C.glfwTerminate()
}

func GetVersion() (int, int, int) {
	var (
		major C.int
		minor C.int
		rev   C.int
	)
	C.glfwGetVersion(&major, &minor, &rev)
	return int(major), int(minor), int(rev)
}

func GetVersionString() string {
	return C.GoString(C.glfwGetVersionString())
}
