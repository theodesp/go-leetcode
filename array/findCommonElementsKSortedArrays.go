package array

/**
Find all common elements in K arrays
for example
1 var arr1 = [1, 5, 5, 10];
2 var arr2 = [3, 4, 5, 5, 10];
3 var arr3 = [5, 5, 10, 20];
4 var output = [5 ,10];

In this example with three arrays, k=3.
 */

/* Solution:

To do this, simply iterate over each array and count instances of every element. However, do not
track repeated ones. Use a temp variable last to keep track of the previous element on each array.
This is to avoid increasing the counter when we see a duplicate element in the same array.

 */
func FindCommonElementsKSortedArrays(arrs [][]int) []int  {
	seen := make(map[int]int)
	var last int
	for i:=0;i<len(arrs);i+=1 {
		for _, el := range arrs[i] {
			if last != el {
				if _, ok := seen[el]; !ok {
					seen[el] = 1
				} else {
					seen[el] += 1
				}
			}
			last = el
		}
	}

	result := []int{}
	for key, _ := range seen {
		if seen[key] == len(arrs) {
			result = append(result, key)
		}
 	}

	return result

}
