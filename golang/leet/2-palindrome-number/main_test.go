package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, isPalindromeSol1(121), true)
	assert.Equal(t, isPalindromeSol1(-121), false)
	assert.Equal(t, isPalindromeSol1(10), false)

	assert.Equal(t, isPalindromeSol2(121), true)
	assert.Equal(t, isPalindromeSol2(-121), false)
	assert.Equal(t, isPalindromeSol2(10), false)
}
