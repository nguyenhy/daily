/*
 - problem: https://leetcode.com/problems/valid-parentheses/
 - solution:
*/
package main

import (
	"fmt"
)

func main() {

	fmt.Println(isValid("[") == false)
	fmt.Println(isValid("[)") == false)
	fmt.Println(isValid("{[]}") == true)
	fmt.Println(isValid("{][}") == false)
	fmt.Println(isValid("{][]}") == false)

}

func isValid(s string) bool {
	rawLength := len(s)
	if rawLength%2 == 1 {
		return false
	}

	open := []string{}
	for i := 0; i < rawLength; i++ {
		char := s[i : i+1]

		if char == "[" || char == "(" || char == "{" {
			open = append(open, char)

		} else {
			openLength := len(open)

			if openLength == 0 {
				return false
			}
			last := open[openLength-1]
			if (char == "]" && last == "[") || (char == ")" && last == "(") || (char == "}" && last == "{") {
				open = open[:openLength-1]
			} else {
				return false
			}
		}
	}

	return len(open) == 0
}
