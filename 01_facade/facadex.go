package facade


type Draw struct {
	circle *Circle
	rectangle *Rectangle
}

func NewDraw() *Draw {
	return &Draw{
		circle: NewCircle(),
		rectangle: NewRectangle(),
	}
}
func (ts *Draw) DrawCircle() string  {
	return "circle"
}
func (ts *Draw) DrawRectangle() string  {
	return "rectangle"
}
//----------------------------------------------
func NewCircle() *Circle {
	return &Circle{}
}
type Circle struct {
	name string
}

func (ts *Circle) draw() string  {
	return "circle"
}
//----------------------------------------------
func NewRectangle() *Rectangle {
	return &Rectangle{}
}
type Rectangle  struct {
	name string
}

func (ts *Rectangle) draw() string  {
	return "rectangle"
}
//----------------------------------------------