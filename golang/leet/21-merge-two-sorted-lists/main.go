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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var pointer *ListNode = list1
	if list1.Val > list2.Val {
		pointer = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}

	var current *ListNode = pointer
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else if list2 == nil {
		current.Next = list1
	}

	return pointer
}
