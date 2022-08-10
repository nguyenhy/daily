/*
 - problem: https://leetcode.com/problems/longest-common-prefix/
 - solution:
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(longestCommonPrefix([]string{"b", ""}) == "")
	fmt.Println(longestCommonPrefix([]string{"a", "ab", ""}) == "")
	fmt.Println(longestCommonPrefix([]string{"", "b"}) == "")
	fmt.Println(longestCommonPrefix([]string{"", "a", "ab"}) == "")
	fmt.Println(longestCommonPrefix([]string{"", "", ""}) == "")
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}) == "fl")
	fmt.Println(longestCommonPrefix([]string{"ab", "ad"}) == "a")
	fmt.Println(longestCommonPrefix([]string{"reflower", "flow", "flight"}) == "")
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}) == "")
	fmt.Println(longestCommonPrefix([]string{"ab", "ab"}) == "ab")
	fmt.Println(longestCommonPrefix([]string{"flower", "fkow"}) == "f")
	fmt.Println("----------------------")
}

func longestCommonPrefix(strs []string) string {
	// fmt.Println("----------------------", strs)

	length := len(strs)

	if length < 1 {
		return ""
	}

	if length == 1 {
		return strs[0]
	}

	var shortest string = ""
	for i := 0; i < length; i++ {
		str := strs[i]

		if len(str) <= 0 {
			return ""
		}

		if len(shortest) <= 0 {
			shortest = str
			continue
		}

		if len(shortest) > len(str) {
			shortest = str
		}
	}

	// fmt.Println("shortest", shortest)
	shortestLength := len(shortest)

	if shortestLength <= 0 {
		return ""
	}

	prefix := ""
	for i := 0; i < shortestLength; i++ {
		str := shortest[0 : i+1]
		if hasPrefix(str, strs) {
			prefix = str
		} else {
			return prefix
		}
	}

	return prefix
}

func hasPrefix(prefix string, strs []string) bool {
	length := len(strs)
	for i := 0; i < length; i++ {
		string := strs[i]
		contain := strings.HasPrefix(string, prefix)
		// fmt.Println("hasPrefix", string, prefix, contain)
		if !contain {
			return false
		}
	}
	return true
}
