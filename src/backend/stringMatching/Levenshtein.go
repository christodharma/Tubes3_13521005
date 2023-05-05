package StringMatching

func MinMulti(nums ...int) int {
    if len(nums) == 0 {
        return 0
    }
    m := nums[0]
    for _, n := range nums[1:] {
        if n < m {
            m = n
        }
    }
    return m
}

func Levenshtein(s, t string) int {
    m, n := len(s), len(t)
    if m == 0 {
        return n
    }
    if n == 0 {
        return m
    }

    // Create a 2D slice to hold the distances
    d := make([][]int, m+1)
    for i := range d {
        d[i] = make([]int, n+1)
    }

    // Initialize the first row and column of the distance matrix
    for i := 0; i <= m; i++ {
        d[i][0] = i
    }
    for j := 0; j <= n; j++ {
        d[0][j] = j
    }

    // Calculate the Levenshtein distance
    for j := 1; j <= n; j++ {
        for i := 1; i <= m; i++ {
            if s[i-1] == t[j-1] {
                d[i][j] = d[i-1][j-1]
            } else {
                d[i][j] = MinMulti(
                    d[i-1][j]+1,   // deletion
                    d[i][j-1]+1,   // insertion
                    d[i-1][j-1]+1, // substitution
                )
            }
        }
    }
	var result int = d[m][n]
	var percentage float64 = float64(result) / float64(n)
    return int(percentage*100)
}