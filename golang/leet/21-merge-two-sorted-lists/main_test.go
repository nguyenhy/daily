package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, (&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}).ToString(), "[1,2,3]")

	assert.Equal(
		t,
		"[0]",
		mergeTwoLists(nil, &ListNode{}).ToString(),
	)
	assert.Equal(
		t,
		"[0]",
		mergeTwoLists(&ListNode{Val: 0}, nil).ToString(),
	)
	assert.Equal(
		t,
		"[0]",
		mergeTwoLists(&ListNode{Val: 0}, nil).ToString(),
	)
	assert.Equal(
		t,
		"[0]",
		mergeTwoLists(&ListNode{Val: 0}, nil).ToString(),
	)
	assert.Equal(
		t,
		"[1,1,2,2,3,3]",
		mergeTwoLists(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		).ToString(),
	)

	assert.Equal(
		t,
		"[1,1,2,3,4,4]",
		mergeTwoLists(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		).ToString(),
	)

	assert.Equal(
		t,
		"[1,2]",
		mergeTwoLists(
			&ListNode{
				Val: 2,
			},
			&ListNode{
				Val: 1,
			},
		).ToString(),
	)
}
