package sorting

func MergeSort(a []int) []int {

	if len(a) < 2 {
		return a
	}
	mid := (len(a)) / 2
	return Merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func Merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
