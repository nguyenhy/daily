// https://leetcode.com/problems/product-of-array-except-self/?envType=study-plan-v2&envId=leetcode-75
package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	var length = len(nums)
	var output = make([]int, length)

	output[0] = 1

	for i := 1; i < length; i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	var right int = 1
	for i := length - 1; i >= 0; i-- {
		output[i] = output[i] * right
		right = right * nums[i]
	}

	return output
}
