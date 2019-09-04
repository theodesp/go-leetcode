package uncategorized

import "math"

/*
Say you have an array for which the ith element is the price of
a given stock on day i.

If you were only permitted to complete at most one transaction
(i.e., buy one and sell one share of the stock),
design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
 */

 /*
 Solution 1
 1) Brute force by having two for loops.
 Set current max=0
 First loop starts with i at 0 index.
 Second loop starts with j at i + 1 index.
 Calculate prices[j] - prices[i]. If > 0 then update max if value bigger that current max otherwise skip
 Return max

 Time Complexity: O(n^2)
 Space Complexity: O(1)
  */
func MaxProfitBrute(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i += 1 {
		for j := i + 1; j < len(prices); j += 1 {
			diff := prices[j] - prices[i]
			if diff > 0 && diff > max {
				max = diff
			}
		}
	}
	return max
}

/*
 Solution 2
 1) Optimised with one loop.
 Set current max=0
 Set min = A really high number
 First loop starts with i at 0 index.
 Check if prices[i] < current min. if yes update min
 Otherwise if min - prices[i] > current max
 Update max
 Return max

 Time Complexity: O(n)
 Space Complexity: O(1)
  */
func MaxProfit(prices []int) int {
	max := int32(0)
	min := math.MaxInt32
	for i := int32(0); i < int32(len(prices)); i += 1 {
		if prices[i] < min {
			min = prices[i]
		} else {
			diff := int32(prices[i] - min)
			if diff > 0 && diff > max {
				max = diff
			}
		}
	}
	return int(max)
}
