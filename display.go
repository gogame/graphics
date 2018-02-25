package graphics

import (
	"log"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

//
type Display struct {
	window *glfw.Window
}

/*
const {
	DisplayProfile2D_8bits
}
*/
//
func NewDisplay(width int, height int, scale int) (d *Display) {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(width*scale, height*scale, "", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	d = new(Display)
	d.window = window
	return
}

func (d *Display) Close() {
	glfw.Terminate()
}
