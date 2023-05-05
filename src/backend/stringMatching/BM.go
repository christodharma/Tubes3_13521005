package StringMatching

func BMMatch(text string, pattern string) int {
    n := len(text)
    m := len(pattern)
    if m == 0 {
        return 0
    }
    if n < m {
        return -1
    }
    skip := make(map[rune]int)
    for k := range pattern {
        skip[rune(pattern[k])] = k
    }
    k := m - 1
    for k < n {
        j := m - 1
        i := k
        for j >= 0 && text[i] == pattern[j] {
            i--
            j--
        }
        if j == -1 {
            return i + 1
        }
        if skip[rune(text[k])] >= 0 {
            k += m - skip[rune(text[k])] - 1
        } else {
            k += m
        }
    }
    return -1
}
