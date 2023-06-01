package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	nums := []int{2, 2, 7, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
