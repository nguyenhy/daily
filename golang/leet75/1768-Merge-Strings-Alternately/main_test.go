// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package mergestringsalternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, mergeAlternately("abc", "pqr"), "apbqcr")
	assert.Equal(t, mergeAlternately("ab", "pqrs"), "apbqrs")
	assert.Equal(t, mergeAlternately("abcd", "pq"), "apbqcd")

	assert.Equal(t, mergeAlternately("a", "p"), "ap")
}
