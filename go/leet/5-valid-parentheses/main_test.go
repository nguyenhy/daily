package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(isValid("[") == false)
	fmt.Println(isValid("[)") == false)
	fmt.Println(isValid("{[]}") == true)
	fmt.Println(isValid("{][}") == false)
	fmt.Println(isValid("{][]}") == false)
}
