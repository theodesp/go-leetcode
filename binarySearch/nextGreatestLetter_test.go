package binarySearch

import (
	"testing"
)

func TestNextGreatestLetter(t *testing.T)  {
	if NextGreatestLetter([]byte("cfj"), byte('a')) != byte('c') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("cfj"), byte('c')) != byte('f') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("cfj"), byte('d')) != byte('f') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("cfj"), byte('g')) != byte('j') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("cfj"), byte('j')) != byte('c') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("cfj"), byte('k')) != byte('c') {
		t.Fail()
	}

	if NextGreatestLetter([]byte("eeeeeennnn"), byte('e')) != byte('n') {
		t.Fail()
	}
}
