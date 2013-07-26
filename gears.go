// +build ignore

package main

/*
 * 3-D gear wheels.
 *
 * Command line options:
 *    -info      print GL implementation information
 *    -exit      automatically exit after 30 seconds
 *
 */

import (
	"flag"
	"fmt"
	"github.com/go-gl/gl"
	glfw "github.com/grd/glfw3"
	"math"
	"os"
	"time"
)

// If non-zero, the program exits after that many seconds
var autoexit *int

/*
  Draw a gear wheel.  You'll probably want to call this function when
  building a display list since we do a lot of trig here.

  Input:  inner_radius - radius of hole at center
          outer_radius - radius at center of teeth
          width - width of gear teeth - number of teeth
          tooth_depth - depth of tooth

*/
func gear(inner_radius, outer_radius, width float64,
	teeth int, tooth_depth float64) {

	var i int
	var angle float64
	var u, v, len float64

	r0 := inner_radius
	r1 := outer_radius - tooth_depth/2.0
	r2 := outer_radius + tooth_depth/2.0

	da := 2.0 * math.Pi / float64(teeth) / 4.0

	gl.ShadeModel(gl.FLAT)

	gl.Normal3d(0.0, 0.0, 1.0)

	// draw front face
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		if i < teeth {
			gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
			gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
		}
	}
	gl.End()

	// draw front sides of teeth
	gl.Begin(gl.QUADS)
	da = 2.0 * math.Pi / float64(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
	}
	gl.End()

	gl.Normal3d(0.0, 0.0, -1.0)

	// draw back face
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		if i < teeth {
			gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
			gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		}
	}
	gl.End()

	// draw back sides of teeth
	gl.Begin(gl.QUADS)
	da = 2.0 * math.Pi / float64(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), -width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), -width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
	}
	gl.End()

	// draw outward faces of teeth
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
		u = r2*math.Cos(angle+da) - r1*math.Cos(angle)
		v = r2*math.Sin(angle+da) - r1*math.Sin(angle)
		len = math.Sqrt(u*u + v*v)
		u /= len
		v /= len
		gl.Normal3d(v, -u, 0.0)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), -width*0.5)
		gl.Normal3d(math.Cos(angle), math.Sin(angle), 0.0)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), -width*0.5)
		u = r1*math.Cos(angle+3*da) - r2*math.Cos(angle+2*da)
		v = r1*math.Sin(angle+3*da) - r2*math.Sin(angle+2*da)
		gl.Normal3d(v, -u, 0.0)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
		gl.Normal3d(math.Cos(angle), math.Sin(angle), 0.0)
	}

	gl.Vertex3d(r1*math.Cos(0), r1*math.Sin(0), width*0.5)
	gl.Vertex3d(r1*math.Cos(0), r1*math.Sin(0), -width*0.5)

	gl.End()

	gl.ShadeModel(gl.SMOOTH)

	// draw inside radius cylinder
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Normal3d(-math.Cos(angle), -math.Sin(angle), 0.0)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
	}
	gl.End()

}

var (
	view_rotx           = 20.0
	view_roty           = 30.0
	view_rotz           = 0.0
	gear1, gear2, gear3 uint
	angle               = 0.0
)

// OpenGL draw function & timing
func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Rotated(view_rotx, 1.0, 0.0, 0.0)
	gl.Rotated(view_roty, 0.0, 1.0, 0.0)
	gl.Rotated(view_rotz, 0.0, 0.0, 1.0)

	gl.PushMatrix()
	gl.Translated(-3.0, -2.0, 0.0)
	gl.Rotated(angle, 0.0, 0.0, 1.0)
	gl.CallList(gear1)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translated(3.1, -2.0, 0.0)
	gl.Rotated(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	gl.CallList(gear2)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translated(-3.1, 4.2, 0.0)
	gl.Rotated(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	gl.CallList(gear3)
	gl.PopMatrix()

	gl.PopMatrix()
}

// update animation parameters
func animate() {
	angle = 100.0 * glfw.GetTime()
}

// change view angle, exit upon ESC
func key(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch glfw.Key(k) {
	case glfw.KeyZ:
		if mods&glfw.ModShift != 0 {
			view_rotz -= 5.0
		} else {
			view_rotz += 5.0
		}
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	case glfw.KeyUp:
		view_rotx += 5.0
	case glfw.KeyDown:
		view_rotx -= 5.0
	case glfw.KeyLeft:
		view_roty += 5.0
	case glfw.KeyRight:
		view_roty -= 5.0
	default:
		return
	}
}

// new window size
func reshape(window *glfw.Window, width, height int) {
	h := float64(height) / float64(width)

	znear := 5.0
	zfar := 30.0
	xmax := znear * 0.5

	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-xmax, xmax, -xmax*h, xmax*h, znear, zfar)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translated(0.0, 0.0, -20.0)
}

// program & OpenGL initialization
func progInit() {
	pos := []float32{5.0, 5.0, 10.0, 0.0}
	red := []float32{0.8, 0.1, 0.0, 1.0}
	green := []float32{0.0, 0.8, 0.2, 1.0}
	blue := []float32{0.2, 0.2, 1.0, 1.0}

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	// make the gears
	gear1 = gl.GenLists(1)
	gl.NewList(gear1, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, red)
	gear(1.0, 4.0, 1.0, 20, 0.7)
	gl.EndList()

	gear2 = gl.GenLists(1)
	gl.NewList(gear2, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, green)
	gear(0.5, 2.0, 2.0, 10, 0.7)
	gl.EndList()

	gear3 = gl.GenLists(1)
	gl.NewList(gear3, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, blue)
	gear(1.3, 2.0, 0.5, 10, 0.7)
	gl.EndList()

	gl.Enable(gl.NORMALIZE)

	// Parse command line options

	info := flag.Bool("info", false, "Info")
	autoexit = flag.Int("exit", 30, "Auto Exit after n seconts\n")

	flag.Parse()

	if *info {
		fmt.Printf("gl.RENDERER   = %s\n", gl.GetString(gl.RENDERER))
		fmt.Printf("gl.VERSION    = %s\n", gl.GetString(gl.VERSION))
		fmt.Printf("gl.VENDOR     = %s\n", gl.GetString(gl.VENDOR))
		fmt.Printf("gl.EXTENSIONS = %s\n", gl.GetString(gl.EXTENSIONS))
		os.Exit(1)
	}
}

func theProgram() {
	time.Sleep(100 * time.Millisecond)

	glfw.WindowHint(glfw.DepthBits, 16)

	window, err := glfw.CreateWindow(300, 300, "Gears", nil, nil)
	if err != nil {
		panic(err)
	}

	// Set callback functions
	window.SetFramebufferSizeCallback(reshape)
	window.SetKeyCallback(key)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	width, height := window.GetFramebufferSize()
	reshape(window, width, height)

	// Parse command-line options
	progInit()

	// Main loop
	for !window.ShouldClose() {
		// Draw gears
		draw()

		// Update animation
		animate()

		// Swap buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}
	os.Exit(0)
}

func main() {
	go theProgram()
	glfw.Main()
}
