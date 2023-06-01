// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
package main

import (
	"testing"
)

func expect(realValue any, expectValue any, t *testing.T) {
	if realValue != expectValue {
		t.Errorf("expect: |%v|, real: |%v|", expectValue, realValue)
	}
}

func Test(t *testing.T) {
	expect(gcd(12, 6), 6, t)

	expect(gcdOfStrings("ABCABC", "ABC"), "ABC", t)
	expect(gcdOfStrings("ABABAB", "ABAB"), "AB", t)
	expect(gcdOfStrings("LEET", "CODE"), "", t)

	expect(gcdOfStrings("a", "a"), "a", t)
}
