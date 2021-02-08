package adapter

import "testing"


func TestAdapter(t *testing.T) {
	ori := NewOriginStruct()
	target := NewAdapter(ori)
	res := target.Request()
	if res != "origin" {
		t.Fatal("adapter test fail")
	}
}
