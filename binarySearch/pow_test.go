package binarySearch

import (
	"testing"
)

func TestMyPow(t *testing.T)  {
	if MyPow(2, 2) != 4 {
		t.Fail()
	}

	if MyPow(2, -2) != 0.25000 {
		t.Fail()
	}
}
