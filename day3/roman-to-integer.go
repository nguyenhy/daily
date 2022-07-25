/*
 - problem: https://leetcode.com/problems/roman-to-integer/
 - solution:
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	length := len(s)

	romanCharMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	fmt.Println(romanCharMap)
	var output int = 0
	for i := length - 1; i >= 0; i-- {
		romanNumber := string(s[i])
		intNumber := romanCharMap[romanNumber]

		fmt.Println(i, romanNumber, intNumber, output)
		if intNumber >= 5 {
			output += intNumber
		} else {
			output -= intNumber
		}
		fmt.Println("->", output)
	}

	return output
}
