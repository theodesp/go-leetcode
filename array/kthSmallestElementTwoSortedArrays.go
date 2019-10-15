package array

/*
Given two sorted arrays of size m and n respectively, 
you are tasked with finding the element that would be at the k’th 
position of the final sorted array.
 */


/* Solution 1:
Since we are given two sorted arrays, we can use merging technique to get the final merged array.
From this, we simply go to the k’th index.

Complexity O(n+m)
 */
func FindKthElementTwoSortedArraysBasic(a, b []int, k int) int  {
	sorted := make([]int, len(a)+len(b))
	i,j,m := 0,0,0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted[m] = a[i]
			i += 1
		} else {
			sorted[m] = b[j]
			j += 1
		}
		m += 1
	}

	for i < len(a) {
		sorted[m] = a[i]
		i += 1
		m += 1
	}

	for j < len(b) {
		sorted[m] = b[j]
		j += 1
		m += 1
	}
	return sorted[k-1]
}

/* Solution 2:
We compare the middle elements of arrays arr1 and arr2,
let us call these indices mid1 and mid2 respectively.

Let us assume arr1[mid1]  k, then clearly the elements after
mid2 cannot be the required element. We then set the last
element of arr2 to be arr2[mid2].

In this way, we define a new subproblem with half the size
of one of the arrays.

Complexity O(log n + log m)
*/
func FindKthElementTwoSortedArrays(a, b []int, k int) int  {
	lo1, hi1 := 0, len(a) - 1
	lo2, hi2 := 0, len(b) - 1

	// if a == empty
	if lo1 == hi1 || len(a) == 0 {
		return b[k]
	}

	// if b == empty
	if lo2 == hi2 || len(b) == 0 {
		return a[k]
	}

	// Find first middle element
	m1 := int(uint(lo1+hi1) >> 1)
	// Find second middle element
	m2 := int(uint(lo2+hi2) >> 1)

	// if k > m1 + m2 means that k exists between a[m1:] and b[m2:].
	// Then if a[m1] > b[m2] search over the entire a and b[m2+1:] for the (k-m2-1)th element
	// As we know than if we merge a[m1:] and b[m2:] then anything from b[:m2] will fit before a[m1:].
	// Otherwise search the entire b and a[m1+1:] for the (k-m1-1)element
	if m1+m2 < k {
		if a[m1] > b[m2] {
			return FindKthElementTwoSortedArrays(a, b[m2+1:], k-m2-1)
		} else {
			return FindKthElementTwoSortedArrays(a[m1:], b, k-m1-1)
		}
	} else {
		// else if k <= m1 + m2 means that k exists between a[:m1] and b[:m2].
		// Then if a[m1] > b[m2] search over the a[0:m1] and all b
		// as we will know that if we sort a[0:m1] and b[:m2] b will fit between a[0:m1].
		// Otherwise we do the same for b as we will know that a will fit between b[0:m2]
		if a[m1] > b[m2] {
			return FindKthElementTwoSortedArrays(a[:m1], b, k)
		} else {
			return FindKthElementTwoSortedArrays(a, b[:m2], k)
		}
	}

}

