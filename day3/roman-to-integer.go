/*
 - problem: https://leetcode.com/problems/roman-to-integer/
 - solution:
*/
package main

import (
	"fmt"
)

func main() {
	test := []string{"III", "II", "IV", "V", "VI", "VII", "VIII", "IX", "XII", "XIII", "XIV", "XIX", "MCMXCIV", "XC"}
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i], romanToInt(test[i]))
		fmt.Println("----------------------")
	}
}

func romanToInt(s string) int {

	// create roman to integer map
	romanCharMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	// check if `s` in `romanCharMap`, for better perf
	if romanCharMap[s] > 0 {
		return romanCharMap[s]
	}

	length := len(s)
	var output int = 0

	// loop from right to left
	for i := length - 1; i >= 0; i-- {
		// roman number at `i` position
		romanNumber := string(s[i])

		// corresponding integer value of `romanNumber`
		intNumber := romanCharMap[romanNumber]

		// only increment if new value is at least `5x` current sum (output)
		if 5*intNumber > output {
			/*
				Eg: XC ( 90 )
				index 		= 1
				intNumber 	= 100
				output	 	= 0
				--------
				5 * 100 > 0 : true
				------
				output 		= 0 + 100


				index 		= 0
				intNumber 	= 10
				output	 	= 100
				--------
				5 * 10 > 100: false
				------
				output 		= 100 - 10 = 90
			*/
			output += intNumber
		} else {
			output -= intNumber
		}
	}

	return output
}
