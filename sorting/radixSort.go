package sorting

/*
We can’t use counting sort because counting sort will take O(n2)
which is worse than comparison based sorting algorithms.
Can we sort such an array in linear time?
Radix Sort is the answer. The idea of Radix Sort is to do digit by digit
sort starting from least significant digit to most significant digit.
Radix sort uses counting sort as a subroutine to sort.
 */

func RadixSort(arr []int) []int  {
	// Find the maximum number to know number of digits
	m := findMax(arr)
	result := make([]int,len(arr), len(arr))

	// Do counting sort for every digit. Note that instead
	// of passing digit number, e is passed. e is 10^i
	// where i is current digit number
	for e := 1; m/e > 0; e*=10 {
		result = CountingSortExp(arr, e)
	}
	return result
}

func CountingSortExp(arr []int, e int) []int  {
	output := make([]int,len(arr), len(arr))
	// Digits are base 10 for now
	count := make([]int, 10, 10)
	// Compute int counts
	for i:=0;i<len(arr);i+=1 {
		count[(arr[i]/e) % 10] +=1
	}
	// Modify the count array such that each element at each index
	// stores the sum of previous counts.
	for i:=1;i<len(count);i+=1 {
		count[i] += count[i-1]
	}

	// Build the output character array
	for i:=0;i<len(arr);i+=1 {
		output[count[(arr[i]/e) % 10] - 1] = arr[i]
		count[(arr[i]/e) % 10] -= 1
	}
	return output
}


func findMax(arr []int) int  {
	var currMax int
	for _,val := range arr {
		if val > currMax {
			currMax = val
		}
	}
	return currMax
}