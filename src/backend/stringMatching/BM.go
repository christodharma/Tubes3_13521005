package StringMatching

import "errors"

func bmMatch(text string, pattern string) (int, error) {
    last := buildLast(pattern)
    n := len(text)
    m := len(pattern)
    i := m - 1

    if i > n-1 {
        return -1, errors.New("pattern is longer than text")
    }

    j := m - 1

    for i <= n-1 {
        if pattern[j] == text[i] {
            if j == 0 {
                return i, nil
            }
            i--
            j--
        } else {
            lo := last[text[i]]
            i += m - min(j, 1+lo)
            j = m - 1
        }
    }

    return -1, nil
}

func buildLast(pattern string) []int {
    last := make([]int, 256) // assume ASCII character set
    m := len(pattern)

    for i := 0; i < 256; i++ {
        last[i] = -1
    }

    for i := 0; i < m; i++ {
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
