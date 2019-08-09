package uncategorized

import "fmt"
// http://btechsmartclass.com/data_structures/knuth-morris-pratt-algorithm.html
func KmpSearch(needle string, heystack string) int {
	m := len(needle)
	n := len(heystack)
	lps := GetLps(heystack)
	fmt.Println(&lps)
	i := 0
	j := 0
	for i < n {
		if needle[j] == heystack[i] {
			i += 1
			j += 1
		}
		if j == m {
			return i - j
		} else if i < n && needle[j] != heystack[i] {
			if j != 0 {
				j = lps[j-1] // Line here uses computed array
			} else {
				i += 1
			}
		}
	}

	return -1
}
