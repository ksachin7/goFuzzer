package main

import "testing"

func FuzzReverse(f *testing.F) {
    testCases := []string{
        "hello", 
        "world", 
        "12345", 
        "!@#$%", 
        "日本語",        // Unicode
        "😊😊😊",       // Emojis
        string(make([]rune, 10000)), // Very large input
    }
    for _, tc := range testCases {
        f.Add(tc)
    }

    f.Fuzz(func(t *testing.T, input string) {
        reversed := Reverse(input)
        doubleReversed := Reverse(reversed)
        if input != doubleReversed {
            t.Errorf("Expected %q but got %q", input, doubleReversed)
        }
    })
}

func FuzzPalindromeCheck(f *testing.F) {
    f.Add("madam")
    f.Add("hello")
    f.Fuzz(func(t *testing.T, input string) {
        isPalindrome := PalindromeCheck(input)
        if len(input) > 0 && input == Reverse(input) && !isPalindrome {
            t.Errorf("PalindromeCheck failed for %q", input)
        }
    })
}
