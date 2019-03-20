package binarySearch

// Floor of square root of x
func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	lo, hi, res := 1, x, -1
	// Do Binary Search for floor(sqrt(x))
	for lo <= hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		// If x is a perfect square
		if m*m == x {
			return m
		}

		if m*m < x {
			lo = m + 1
			res = m
		} else {
			hi = m
		}
	}
	return res
}
