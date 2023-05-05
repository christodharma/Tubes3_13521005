package StringMatching
// package main

// import ("fmt")

func KMPMatch(pattern string, text string) []int {
    m := len(pattern)
    n := len(text)

    // Compute the longest prefix suffix array for the pattern
    lps := make([]int, m)
    j := 0
    for i := 1; i < m; i++ {
        for j > 0 && pattern[j] != pattern[i] {
            j = lps[j-1]
        }
        if pattern[j] == pattern[i] {
            j++
        }
        lps[i] = j
    }

    // Use the lps array to perform the pattern search
    result := []int{}
    i := 0
    j = 0
    for i < n {
        if pattern[j] == text[i] {
            i++
            j++
        }
        if j == m {
            result = append(result, i-j)
            j = lps[j-1]
        } else if i < n && pattern[j] != text[i] {
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }

    return result
}

// func main(){
//     pattern := "example"
//     text := "This is an example text"
//     result := KMPMatch(pattern, text)

//     if len(result) > 0 {
//         fmt.Printf("Pattern found at index/indices: %v\n", result)
//     } else {
//         fmt.Println("Pattern not found")
//     }
// }

