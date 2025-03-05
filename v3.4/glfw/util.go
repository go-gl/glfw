package glfw

//#include <stdlib.h>
//#define GLFW_INCLUDE_NONE
//#include "glfw/include/GLFW/glfw3.h"
import "C"

import (
	"image"
	"image/draw"
)

func glfwbool(b C.int) bool {
	return b == C.int(True)
}

func bytes(origin []byte) (pointer *uint8, free func()) {
	n := len(origin)
	if n == 0 {
		return nil, func() {}
	}

	ptr := C.CBytes(origin)
	return (*uint8)(ptr), func() { C.free(ptr) }
}

// imageToGLFW converts img to be compatible with C.GLFWimage.
func imageToGLFW(img image.Image) (r C.GLFWimage, free func()) {
	b := img.Bounds()

	r.width = C.int(b.Dx())
	r.height = C.int(b.Dy())

	var pixels []byte
	if m, ok := img.(*image.NRGBA); ok && m.Stride == b.Dx()*4 {
		pixels = m.Pix[:m.PixOffset(m.Rect.Min.X, m.Rect.Max.Y)]
	} else {
		m := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), img, b.Min, draw.Src)
		pixels = m.Pix
	}

	pix, free := bytes(pixels)
	r.pixels = (*C.uchar)(pix)
	return r, free
}
