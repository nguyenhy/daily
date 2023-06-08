// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package canplaceflowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
	assert.Equal(t, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
	assert.Equal(t, canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2), false)
}
