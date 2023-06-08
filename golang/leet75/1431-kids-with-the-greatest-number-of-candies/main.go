// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&envId=leetcode-75
package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for i := 0; i < len(candies); i++ {
		var item = candies[i]
		if item > max {
			max = item
		}
	}

	var output []bool
	for i := 0; i < len(candies); i++ {
		var item = candies[i]
		output = append(output, item+extraCandies >= max)
	}

	return output
}
