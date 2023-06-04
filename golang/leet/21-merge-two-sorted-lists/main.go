/*
- problem: https://leetcode.com/problems/merge-two-sorted-lists/
- solution:
*/
package mergetwosortedlists

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) ToString() string {
	if node == nil {
		return "[]"
	}

	result := "["

	for node != nil {
		result += strconv.Itoa(node.Val)

		if node.Next != nil {
			result += ","
		}

		node = node.Next
	}

	result += "]"

	return result
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) /* *ListNode */ {
}
