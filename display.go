package graphics

//
type Display struct {
	i int
}

//
func NewDisplay(width int, height int) *Display {
	return new(Display)
}

func (d *Display) Close() {

}
