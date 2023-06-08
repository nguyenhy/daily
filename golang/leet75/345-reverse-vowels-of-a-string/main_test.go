package reversevowelsofastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, reverseVowels("hello"), "holle")
	assert.Equal(t, reverseVowels("leetcode"), "leotcede")
	assert.Equal(t, reverseVowels("aeiou"), "uoiea")
	assert.Equal(t, reverseVowels("aA"), "Aa")

}
