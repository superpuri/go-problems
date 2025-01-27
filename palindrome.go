package main

import "fmt"

func isPalindrome(str string) bool {
    for i, j := len(str)-1, 0; j <= len(str)/2; i, j = i-1, j+1 {
        if str[i] != str[j] { return false }
    }
    return true
}

func main() {
    var str string

    fmt.Scan(&str)

	if isPalindrome(str) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
