package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, romanToInt("III"), 3)
	assert.Equal(t, romanToInt("II"), 2)
	assert.Equal(t, romanToInt("IV"), 4)
	assert.Equal(t, romanToInt("V"), 5)
	assert.Equal(t, romanToInt("VI"), 6)
	assert.Equal(t, romanToInt("VII"), 7)
	assert.Equal(t, romanToInt("VIII"), 8)
	assert.Equal(t, romanToInt("IX"), 9)
	assert.Equal(t, romanToInt("XII"), 12)
	assert.Equal(t, romanToInt("XIII"), 13)
	assert.Equal(t, romanToInt("XIV"), 14)
	assert.Equal(t, romanToInt("XIX"), 19)
	assert.Equal(t, romanToInt("MCMXCIV"), 1994)
	assert.Equal(t, romanToInt("XC"), 90)
}
