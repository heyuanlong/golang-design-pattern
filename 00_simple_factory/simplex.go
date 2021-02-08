package simplefactory

type Draw interface {
	draw() string
}

func NewObj(name string) Draw  {
	switch name {
	case "circle":return &circle{}
	case "rectangle":return &rectangle{}
	}

	return &circle{}
}
//----------------------------------------------
type circle struct {
	name string
}

func (ts *circle) draw() string  {
	return "circle"
}
//----------------------------------------------
type rectangle  struct {
	name string
}

func (ts *rectangle) draw() string  {
	return "rectangle"
}
//----------------------------------------------

