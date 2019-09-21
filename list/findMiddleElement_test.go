package list

import "testing"

func TestList_Middle(t *testing.T) {
	l := List{head: &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: &Node{data: 4, next: nil}}}}}
	middle := l.Middle()
	if middle == 2 {
		t.Fail()
	}
}
