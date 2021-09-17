/*
 * @Author: bill
 * @Date: 2021-09-17 09:05:22
 * @LastEditors: bill
 * @LastEditTime: 2021-09-17 10:00:12
 * @Description:
 * @FilePath: /leetcode-go/structures/listnode.go
 */
package structures

import "fmt"

// ListNode 链表节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Ints(head *ListNode) []int {
	// 链条长度限制，链条深度超出此限制，会 panic
	limit := 100
	lengths := 0
	res := []int{}

	for head != nil {
		lengths++
		if lengths > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 List2Ints 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	p := l
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return l.Next
}

// GetNodeWithVal return the first node with val
func (l *ListNode) GetNodeWithVal(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node head's index is 0
// if pos = -1, no cycle
func Int2ListWhitCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
