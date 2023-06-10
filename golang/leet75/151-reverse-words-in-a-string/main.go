// https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75
package reversewordsinastring

import "strings"

func reverseWords1(s string) string {
	var builder strings.Builder
	var last string
	var spaceIndex int = -1
	length := len(s)

	for i := 0; i < length; i++ {
		var index = length - (1 + i)
		var item = string(s[index])

		if item == " " {
			if last == " " {
				// last item is space and this item is also space
				continue
			} else {
				last = item
				builder.WriteString(item)
				spaceIndex = builder.Len() - 1
			}
		} else {
			last = item
			left := builder.String()[:spaceIndex+1]
			right := builder.String()[spaceIndex+1:]
			builder.Reset()
			builder.WriteString(left + item + right)
		}
	}

	return strings.Trim(builder.String(), " ")
}

func reverseWords(s string) string {
	words := strings.Fields(s) // Split the input string into words

	reverse(words) // Reverse the order of the words

	return strings.Join(words, " ") // Join the reversed words with a single space
}

func reverse(words []string) {
	left := 0
	right := len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
}
