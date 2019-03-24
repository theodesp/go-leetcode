package binarySearch

import (
	"testing"
)

func TestPerfectSquare(t *testing.T)  {
	if IsPerfectSquare(16) != true {
		t.Fail()
	}

	if IsPerfectSquare(14) == true {
		t.Fail()
	}
}
