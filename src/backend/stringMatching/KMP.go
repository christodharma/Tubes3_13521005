package StringMatching
// package main

// import ("fmt")

func KMPMatch(text string, pattern string) int {
    n := len(text)
    m := len(pattern)
    if m == 0 {
        return -1
    }
    fail := computeFail(pattern)
    i := 0
    j := 0
    for i < n {
        if pattern[j] == text[i] {
            if j == m-1 {
                return i - m + 1 // match
            }
            i++
            j++
        } else if j > 0 {
            j = fail[j-1]
        } else {
            i++
        }
    }
    return -1 // no match
}

func computeFail(pattern string) []int {
    m := len(pattern)
    if m == 0 {
        return []int{}
    }
    fail := make([]int, m)
    fail[0] = 0
    j := 0
    i := 1
    for i < m {
        if pattern[j] == pattern[i] {
            fail[i] = j + 1 // j+1 chars match
            i++
            j++
        } else if j > 0 { // j follows matching prefix
            j = fail[j-1]
        } else { // no match
            fail[i] = 0
            i++
        }
    }
    return fail
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

