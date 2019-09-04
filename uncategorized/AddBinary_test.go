package uncategorized

import "testing"

func Test_AddBinary(t *testing.T) {

	if AddBinary("11", "1") != "100" {
		t.Fail()
	}

	if AddBinary("1010", "1011") != "10101" {
		t.Fail()
	}
}
