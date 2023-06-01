package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	test := []string{"III", "II", "IV", "V", "VI", "VII", "VIII", "IX", "XII", "XIII", "XIV", "XIX", "MCMXCIV", "XC"}
	for i := 0; i < len(test); i++ {
		fmt.Println(test[i], romanToInt(test[i]))
	}
}
