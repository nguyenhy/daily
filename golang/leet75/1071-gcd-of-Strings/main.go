// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package gcdofstrings

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	length := gcd(len(str1), len(str2))
	return str1[:length]
}
