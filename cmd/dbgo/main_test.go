package main

import "testing"

func TestCalc(t *testing.T) {
	r := Calc()
	if r != 2 {
		t.Error("Error")
	}
}
