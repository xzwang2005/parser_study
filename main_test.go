package main

import (
	"strconv"
	"testing"
)

func TestSingleDigitPlus(t *testing.T) {
	res := Calculate("2+9")
	if res != 11 {
		t.Error("result does not match")
	}
}

func TestSingleDigitMinus(t *testing.T) {
	res := Calculate("9-7")
	if res != 2 {
		t.Error("9-7 != 2")
	}
}

func TestDigitsPlus(t *testing.T) {
	res := Calculate("11 + 22")
	if res != 33 {
		t.Error("11 + 22 != 33")
	}
}

func TestDigitsMinus(t *testing.T) {
	res := Calculate("99 -  57")
	if res != 42 {
		t.Error("99-57!=42")
	}
}

func TestMultipleOperators(t *testing.T) {
	res := Calculate("2 + 6- 4 +5-3")
	if res != 6 {
		t.Error("2 + 6- 4 +5-3!= 6")
	}
}

func TestFourOperations(t *testing.T) {
	res := Calculate("2 + 6* 2 -5/2 +2*3/4+6*8/7")
	if res != 19 {
		t.Error("2 + 6* 2 -5/2 +2*3/4+6*8/7 = ", strconv.Itoa(res))
	}
}
