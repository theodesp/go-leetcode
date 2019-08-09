package uncategorized

import "testing"

func TestDefangingIP(t * testing.T)  {
	if defangIPaddr("1.1.1.1") != "1[.]1[.]1[.]1" {
		t.Fail()
	}

	if defangIPaddr("255.100.50.0") != "255[.]100[.]50[.]0" {
		t.Fail()
	}
}