package StringMatching

func kmpMatch(text, pattern string) int {
    textLength, patternLength := len(text), len(pattern)
    fail := computeFail(pattern)
    i, j := 0, 0
    for i < textLength && j < patternLength {
        if pattern[j] == text[i] {
            i++
            j++
        } else if j > 0 {
            j = fail[j-1]
        } else {
            i++
        }
    }
    if j == patternLength {
        return i - j
    }
    return -1
}
// border function
func computeFail(pattern string) []int {
    n := len(pattern)
    fail := make([]int, n)
    for i, j := 1, 0; i < n; i++ {
        for j > 0 && pattern[i] != pattern[j] {
            j = fail[j-1]
        }
        if pattern[i] == pattern[j] {
            j++
        }
        fail[i] = j
    }
    return fail
}
