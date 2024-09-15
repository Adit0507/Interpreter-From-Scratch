package object

import "testing"

func TestStringHashKey(t *testing.T){
	h1 := &String{Value: "Hello World"}
	h2 := &String{Value: "Hello World"}
	d1 := &String{Value: "My name is johnny"}
	d2 := &String{Value: "My name is johnny"}

	if h1.HashKey() != h2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
		}
		if d1.HashKey() != d2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
		}
		if h1.HashKey() == d1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
		}
}