package simplefactory

import "testing"

//TestType1 test get hiapi with factory
func TestObj(t *testing.T) {
	var obj Draw
	obj = NewObj("circle")
	if obj.draw()  != "circle"{
		t.Fatal("circle test fail")
	}

	obj = NewObj("rectangle")
	if obj.draw()  != "rectangle"{
		t.Fatal("rectangle test fail")
	}

}
