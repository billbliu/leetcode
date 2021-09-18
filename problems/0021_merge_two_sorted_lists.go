/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-18 11:59:05
 * @Description: go test -v  0021_merge_two_sorted_lists_test.go 0021_merge_two_sorted_lists.go types.go
 * @FilePath: /leetcode-go/problems/0021_merge_two_sorted_lists.go
 */
package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}

// 迭代
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	resNode := &ListNode{}
	prev := resNode
	for {
		if l1 == nil {
			prev.Next = l2
			break
		}
		if l2 == nil {
			prev.Next = l1
			break
		}
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
			prev = prev.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
			prev = prev.Next
		}
	}
	return resNode.Next
}
