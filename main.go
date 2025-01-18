package main

import (
    "fmt"
)

func Reverse(input string) string {
    result := []rune(input)
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return string(result)
}
func PalindromeCheck(input string) bool {
    reversed := Reverse(input)
    return input == reversed
}

func main() {
    fmt.Println(Reverse("hello"))
}
