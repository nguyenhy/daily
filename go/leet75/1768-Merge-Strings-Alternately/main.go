// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var minLength int
	if len(word1) > len(word2) {
		minLength = len(word2)
	} else {
		minLength = len(word1)
	}

	var builder strings.Builder
	for i := 0; i < minLength; i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	if len(word1) != len(word2) {
		if len(word1) > len(word2) {
			builder.WriteString(word1[minLength:])
		} else {
			builder.WriteString(word2[minLength:])
		}
	}

	return builder.String()
}
