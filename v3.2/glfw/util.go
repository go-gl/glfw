package glfw

//#include <stdlib.h>
//#include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"reflect"
	"unsafe"
)

func glfwbool(b C.int) bool {
	if b == C.GL_TRUE {
		return true
	}
	return false
}

func bytes(origin []byte) (pointer *uint8, free func()) {
	if len(origin) == 0 {
		panic("Slice of bytes cannot be empty.")
	}

	l := len(origin)
	d := C.malloc(C.size_t(l))

	ds := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(d),
		Len:  l,
		Cap:  l,
	}))

	copy(ds[0:l], origin[:])

	return (*uint8)(unsafe.Pointer(&ds[0])), func() { C.free(d) }
}
