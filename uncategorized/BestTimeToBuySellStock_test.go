package uncategorized

import "testing"

func TestMaxProfitBrute(t * testing.T)  {
	if MaxProfitBrute([]int{7,1,5,3,6,4}) != 5 {
		t.Fail()
	}

	if MaxProfitBrute([]int{7,6,4,3,1}) != 0 {
		t.Fail()
	}
}

func TestMaxProfit(t * testing.T)  {
	if MaxProfit([]int{7,1,5,3,6,4}) != 5 {
		t.Fail()
	}

	if MaxProfit([]int{7,6,4,3,1}) != 0 {
		t.Fail()
	}
}
