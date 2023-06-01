/*
- problem: https://leetcode.com/problems/two-sum
- solution: https://www.code-recipe.com/post/two-sum
*/
package twosum

func twoSum(nums []int, target int) []int {
	// create map of [value:index]
	seen := make(map[int]int)
	for currentIndex, currentValue := range nums {
		// expect this value present in "sums"
		var expect int = target - currentValue

		// check if expect value in nums
		seenIndex, present := seen[expect]
		if present {
			// if true then we found a "expect value" that match with "currentValue" to create "target"
			return []int{seenIndex, currentIndex}
		} else {
			// if false then we
			// - add pair of "currentValue"-"currentIndex" to "seen" map
			// - continue loop to the end of "nums" to check if there's matching value with "currentValue"
			seen[currentValue] = currentIndex
		}

	}

	// there's not matching
	return []int{}
}
