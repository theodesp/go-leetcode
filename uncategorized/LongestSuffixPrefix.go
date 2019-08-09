package uncategorized

func LongestSuffixPrefix(s string) int  {
	n := len(s)
	len := 0
	lps := make([]int, n, n)
	i := 1
	for i < n {
		if s[i] == s[len] {
			len += 1
			lps[i] = len
			i += 1
		} else {
			// If they don't match then they have either
			// no longest prefix and suffix
			if len != 0 {
				len = lps[len-1]
			} else {
				lps[i] = 0
				i += 1
			}
		}
	}

	return len
}

