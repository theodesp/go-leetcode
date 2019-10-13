package sorting

/*
Counting sort is a sorting technique based on keys between a
specific range. It works by counting the number of objects
having distinct key values (kind of hashing).
Then doing some arithmetic to calculate the position of
each object in the output sequence.
 */

func CountingSort(arr []int) []int  {
	// Find lengt of result array
	min, max := findMinMax(arr)
	r := max - min + 1
	count := make([]int, r,r)
	output := make([]int, len(arr),len(arr))
	// Compute int counts
	for i:=0;i<len(arr);i+=1 {
		count[arr[i]] +=1
	}
	// Modify the count array such that each element at each index
	// stores the sum of previous counts.
	for i:=1;i<len(count);i+=1 {
		count[i] += count[i-1]
	}

	// Build the output character array
	for i:=0;i<len(arr);i+=1 {
		output[count[arr[i]] - 1] = arr[i]
		count[arr[i]] -= 1
	}
	return output
}

func findMinMax(arr []int) (int, int)  {
	var currMax, currMin int
	for _,val := range arr {
		if val > currMax {
			currMax = val
		}
		if val < currMin {
			currMin = val
		}
	}
	return currMin, currMax
}