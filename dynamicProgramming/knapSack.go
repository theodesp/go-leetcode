package dynamicProgramming

import (
	"fmt"
	"math"
)

/*
Given weights and values of n items, put these items in a knapsack of capacity W to get the maximum total value in
the knapsack. In other words, given two integer arrays val[0..n-1] and wt[0..n-1] which represent values and weights
associated with n items respectively. Also given an integer W which represents knapsack capacity,
find out the maximum value subset of val[] such that sum of the weights of this subset is smaller than or equal to W.
You cannot break an item, either pick the complete item, or don’t pick it (0-1 property).
 */

/* Solution using DP. We consider two cases. We have a function KnapSack(w,n,wt,val)
 First case we include the item in the knapsack. So its value is v[i] now the problem reduces to
KnapSack(w-v[i],n-1,wt,val). The other case is to not include the item in the KnapSack. So the problem
is KnapSack(w,n-1,wt,val) so the available weight is the same but the no of items is less than one.
So for each case we want to find the maximum of either two expressions.

Base case when n=0 or w=0 then we return 0.
Another case is when the available weight is less than the current item.
In that case we ignore it.
 */

func KnapSack(w, n int, wt []int, vals[]int) int  {
	if n == 0 || w == 0 {
		return 0
	}

	// If weight of the nth item is more
	// than Knapsack capacity W, then
	// this item cannot be included
	// in the optimal solution
	if w < wt[n-1] {
		return KnapSack(w,n-1,wt,vals)
	}
	// Return the maximum of two cases:
	// (1) nth item included
	// (2) not included
	maxIncluded := vals[n-1] + KnapSack(w-wt[n-1],n-1,wt,vals)
	maxNotIncluded := KnapSack(w,n-1,wt,vals)

	return int(math.Max(float64(maxIncluded), float64(maxNotIncluded)))
}


// Same problem using DP tabulation
func KnapSackDP(w, n int, wt []int, vals[]int) int  {
	i, j := 0, 0
	k := make([][]int, n+1)
	// Make n x w table to store intermediate results
	// Compute all values of weights for item 1.
	// Then for item 2 etc. If we keep track of the previous value and the current one.
	// We push only the ones that are higher.
	for i, _ := range k {
		k[i] = make([]int, w+1)
	}

	for i = 0; i <= n; i += 1 {
		for j = 0; j <= w; j += 1 {
			if i == 0 || j == 0 {
				k[i][j] = 0
			} else if wt[i-1] <= j {
				// The max value of not included is already saved
				maxNotIncluded := k[i-1][j]
				// The value of included is looked up before
				maxIncluded := vals[i-1] + k[i-1][j-wt[i-1]]
				k[i][j] = int(math.Max(float64(maxNotIncluded),  float64(maxIncluded)))
			} else {
				// We copy the previous i weight
				k[i][j] = k[i-1][j]
			}
		}
	}
	fmt.Println(k)
	return k[n][j-1]
}
