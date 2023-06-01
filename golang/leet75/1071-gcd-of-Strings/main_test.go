// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package gcdofstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, gcd(12, 6), 6)

	assert.Equal(t, gcdOfStrings("ABCABC", "ABC"), "ABC")
	assert.Equal(t, gcdOfStrings("ABABAB", "ABAB"), "AB")
	assert.Equal(t, gcdOfStrings("LEET", "CODE"), "")

	assert.Equal(t, gcdOfStrings("a", "a"), "a")
}
