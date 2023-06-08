// https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75
package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	var length = len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		var emptyLeft = i == 0 || flowerbed[i-1] == 0
		if !emptyLeft {
			continue
		}

		var emptyRight = i == length-1 || flowerbed[i+1] == 0
		if !emptyRight {
			continue
		}

		count += 1
		flowerbed[i] = 1

	}

	return count >= n
}
