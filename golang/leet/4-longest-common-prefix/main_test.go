package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, longestCommonPrefix([]string{"b", ""}), "")
	assert.Equal(t, longestCommonPrefix([]string{"a", "ab", ""}), "")
	assert.Equal(t, longestCommonPrefix([]string{"", "b"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"", "a", "ab"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"", "", ""}), "")
	assert.Equal(t, longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	assert.Equal(t, longestCommonPrefix([]string{"ab", "ad"}), "a")
	assert.Equal(t, longestCommonPrefix([]string{"reflower", "flow", "flight"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"ab", "ab"}), "ab")
	assert.Equal(t, longestCommonPrefix([]string{"flower", "fkow"}), "f")
}
