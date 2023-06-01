package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
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
