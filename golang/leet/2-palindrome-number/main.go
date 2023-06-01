/*
- problem: https://leetcode.com/problems/palindrome-number/
- solution: https://www.code-recipe.com/post/palindrome-number
*/
package palindromenumber

import (
	"strconv"
)

func isPalindromeSol1(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var output int
	temp := x
	for temp > 0 {
		output = output*10 + temp%10
		temp = temp / 10
	}

	return output == x
}

func isPalindromeSol2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var xStr string = strconv.Itoa(x)
	var length = len(xStr)
	var halfLength = length / 2

	for i := 0; i < halfLength; i++ {
		if xStr[i] != xStr[length-1-i] {
			return false
		}
	}

	return true
}
