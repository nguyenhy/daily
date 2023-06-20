// https://leetcode.com/problems/increasing-triplet-subsequence/?envType=study-plan-v2&envId=leetcode-75
package increasingtripletsubsequence

import "math"

func increasingTriplet(nums []int) bool {
	var length = len(nums)
	if length < 3 {
		return false
	}

	var first int = math.MaxInt
	var second int = math.MaxInt
	for i := 0; i < length; i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
