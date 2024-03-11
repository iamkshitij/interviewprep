package main

import (
	"fmt"
	"strings"
)

func main() {

	str := []string{"RACEACAR", "A", "ABCDEFGFEDCBA", "ABC", "ABCBA", "ABBA", "RACEACAR"}
	for i, s := range str {
		fmt.Printf("Test Case # %d\n", i+1)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		fmt.Printf("The input string is '%s' and the length of the string is %d.\n", s, len(s))
		fmt.Printf("\nIs it a palindrome?.....%v\n", isPalindrome(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}

}

func isPalindrome1(str string) bool {

	lt := 0
	rt := len(str) - 1

	for lt < rt {
		if str[lt] == str[rt] {
			lt++
			rt--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome(str string) bool {

	fmt.Println("Input string: ", str)

	left := 0
	right := len(str) - 1

	for left < right {

		if str[left] != str[right] {
			return false
		}

		right--
		left++

	}

	return true

}
