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
	expect(mergeAlternately("abc", "pqr"), "apbqcr", t)
	expect(mergeAlternately("ab", "pqrs"), "apbqrs", t)
	expect(mergeAlternately("abcd", "pq"), "apbqcd", t)

	expect(mergeAlternately("a", "p"), "ap", t)
}
