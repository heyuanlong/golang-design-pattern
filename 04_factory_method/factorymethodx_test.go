package factorymethod

import "testing"

func compute(factory AbstractFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Show()
}

func TestOperator(t *testing.T) {
	var (
		factory AbstractFactory
	)

	factory = Product1Factory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = Product2Factory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
