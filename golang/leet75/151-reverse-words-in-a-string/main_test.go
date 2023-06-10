// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package reversewordsinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	assert.Equal(t, reverseWords1("the sky is blue"), "blue is sky the")
	assert.Equal(t, reverseWords1("a good   example"), "example good a")
	assert.Equal(t, reverseWords1("  hello world  "), "world hello")
	assert.Equal(t, reverseWords1("   "), "")

	assert.Equal(t, reverseWords("the sky is blue"), "blue is sky the")
	assert.Equal(t, reverseWords("a good   example"), "example good a")
	assert.Equal(t, reverseWords("  hello world  "), "world hello")
	assert.Equal(t, reverseWords("   "), "")
}
