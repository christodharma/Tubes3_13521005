package StringMatching

func BMMatch(text, pattern string) int {
	last := buildLast(pattern)
	n, m := len(text), len(pattern)
	if m > n || m == 0{
		return -1 // no match if pattern is longer than text
	}
	i, j := m-1, m-1
	for i < n {
		if pattern[j] == text[i] {
			if j == 0 {
				return i // match
			}
			i--
			j--
		} else {
			lo := last[text[i]] // last occ
			i += m - min(j, 1+lo)
			j = m - 1
		}
	}
	return -1 // no match
}

func buildLast(pattern string) [128]int {
    if len(pattern) == 0 {
        return [128]int{}
    }
	var last [128]int // ASCII char set
	for i := range last {
		last[i] = -1 // initialize array
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
