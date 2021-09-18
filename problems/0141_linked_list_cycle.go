/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:14
 * @LastEditors: bill
 * @LastEditTime: 2021-09-18 13:44:59
 * @Description: go test -v  0141_linked_list_cycle_test.go 0141_linked_list_cycle.go types.go
 * @FilePath: /leetcode-go/problems/0141_linked_list_cycle.go
 */

package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法一：哈希表
func hasCycle1(head *ListNode) bool {
	hashMap := map[*ListNode]bool{}
	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}
	return false
}

// 方法二：快慢指针
func hasCycle2(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
