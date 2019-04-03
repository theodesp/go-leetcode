package list

import (
	"reflect"
	"testing"
)

func TestList_Reverse(t *testing.T) {
	l := NewList()
	l.Push(20)
	l.Push(4)
	l.Push(15)
	l.Push(85)
	l.Push(11)

	var got []interface{}

	l.ForEach(func(data interface{}) {
		got = append(got, data)
	})
	if !reflect.DeepEqual(got, []interface{}{11, 85, 15, 4, 20}) {
		t.Fail()
	}
	got = []interface{}{}
	l.Reverse()
	l.ForEach(func(data interface{}) {
		got = append(got, data)
	})
	if !reflect.DeepEqual(got, []interface{}{20, 4, 15, 85, 11}) {
		t.Fail()
	}
}
