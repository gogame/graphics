package graphics

import "fmt"

//import "github.com/davecgh/go-spew/spew"

type Device struct {
}

type Texture struct {
}

type Color struct {
	R, G, B, A uint8
}

type Vector struct {
	X, Y int
}

type Rectangle struct {
	X, Y, Width, Height int
}

type SpriteBatch struct {
}

func (sb *SpriteBatch) Begin() {

}

func (sb *SpriteBatch) End() {

}

type DrawOptions struct {
	tint     Color
	size     Vector
	src      Rectangle
	origin   Vector
	rotation float32
}

func Tint(c Color) func(*DrawOptions) {
	return func(o *DrawOptions) {
		o.tint = c
	}
}

func (s SpriteBatch) Draw(t Texture, position Vector, options ...func(*DrawOptions)) {
	o := DrawOptions{tint: Color{255, 255, 255, 255}}
	for _, opt := range options {
		opt(&o)
	}
	//spew.Dump(o)
	fmt.Printf("%+v %v %T\n", o, o, o)
}
