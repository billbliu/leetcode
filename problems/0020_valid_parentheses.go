/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-22 14:46:27
 * @Description: go test -v  0020_valid_parentheses_test.go 0020_valid_parentheses.go
 * @FilePath: /leetcode-go/problems/0020_valid_parentheses.go
 */
package problems

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	hashMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 存放左括号的栈
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		// 当前s[i]是右括号
		if v, ok := hashMap[s[i]]; ok {
			// 栈为空或者右括号对应的左括号与栈顶元素不想等
			if len(stack) == 0 || v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
