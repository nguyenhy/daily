// https://leetcode.com/problems/reverse-vowels-of-a-string/?envType=study-plan-v2&envId=leetcode-75
package reversevowelsofastring

func reverseVowels(s string) string {

	var vowels = map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	length := len(s)
	word := []byte(s)
	for left, right := 0, length-1; left < right; {
		if vowels[string(s[left])] && vowels[string(s[right])] {
			word[left], word[right] = word[right], word[left]
			left++
			right--
		}

		if !vowels[string(s[left])] {
			left++
		}

		if !vowels[string(s[right])] {
			right--
		}
	}

	return string(word)
}
