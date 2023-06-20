// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package increasingtripletsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	assert.Equal(t, increasingTriplet([]int{1, 2, 3, 4, 5}), true)
	assert.Equal(t, increasingTriplet([]int{5, 4, 3, 2, 1}), false)
	assert.Equal(t, increasingTriplet([]int{2, 1, 5, 0, 4, 6}), true)

	assert.Equal(t, increasingTriplet([]int{0}), false)
	assert.Equal(t, increasingTriplet([]int{0, 2, 1}), false)
	assert.Equal(t, increasingTriplet([]int{0, 0, 1}), false)

	assert.Equal(t, increasingTriplet([]int{20, 100, 10, 12, 5, 13}), true)

}
