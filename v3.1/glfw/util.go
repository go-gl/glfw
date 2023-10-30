package glfw

//#include <stdlib.h>
//#define GLFW_INCLUDE_NONE
//#include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"image"
	"image/draw"
	"reflect"
	"unsafe"
	"runtime"
)

func glfwbool(b C.int) bool {
	if b == C.int(True) {
		return true
	}
	return false
}

func bytes(origin []byte) (pointer *uint8, free func()) {
	n := len(origin)

	if n == 0 {
		return nil, func() {}
	}

	data := C.malloc(C.size_t(n))

	dataSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(data),
		Len:  n,
		Cap:  n,
	}))

	copy(dataSlice, origin)

	return &dataSlice[0], func() { C.free(data) }
}

// imageToGLFW converts img to be compatible with C.GLFWimage.
// It may reference the underlying pixels buffer in img.
func imageToGLFW(img image.Image) (r C.GLFWimage, free func()) {
	b := img.Bounds()

	r.width = C.int(b.Dx())
	r.height = C.int(b.Dy())

	if m, ok := img.(*image.NRGBA); ok && m.Stride == b.Dx()*4 {
		r.pixels = (*C.uchar)(&m.Pix[0])
		return r, func() { runtime.KeepAlive(m) }
	}

	m := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)

	pix, free := bytes(m.Pix)
	r.pixels = (*C.uchar)(pix)
	return r, free
}
