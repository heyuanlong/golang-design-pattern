package facade

import "testing"


func TestFacade(t *testing.T) {
	obj := NewDraw()
	if obj.DrawCircle() != "circle"{
		t.Fatal("circle test fail")
	}
	if obj.DrawRectangle() != "rectangle"{
		t.Fatal("rectangle test fail")
	}
}
