package main

import "testing"

func TestPlus(t *testing.T) {
	res := Calculate("2+9")
	if res != 11 {
		t.Error("result does not match")
	}
}
