package dynamicProgramming

// Dynamic programming using cache. Bottom down approach
func FibDynamic(n int) int  {
	memory := make(map[int]int, n+1)

	return FibDynamicRecur(memory, n)
}

func FibDynamicRecur(memory map[int]int, n int) int  {
	if n < 2 {
		return n
	}

	if val,ok := memory[n]; ok {
		return val
	}
	memory[n] = FibDynamicRecur(memory, n-1) + FibDynamicRecur(memory, n-2)
	return memory[n]
}

// Dynamic programming using table. Bottom up approach
func FibDynamicUp(n int) int  {
	dp := make(map[int]int, n+1)
	//base cases
	dp[0] = 0
	dp[1] = 1
	for i := 2;i<=n;i+=1 {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}