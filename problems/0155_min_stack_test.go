/*
 * @Author: bill
 * @Date: 2021-09-22 17:45:31
 * @LastEditors: bill
 * @LastEditTime: 2021-09-23 09:06:35
 * @Description: go test -v  0155_min_stack_test.go 0155_min_stack.go
 * @FilePath: /leetcode-go/problems/0155_min_stack_test.go
 */

package problems

import (
	"fmt"
	"testing"
)

func Test_Problem155(t *testing.T) {
	obj1 := Constructor155()
	obj1.Push(1)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Push(0)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Push(10)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Pop()
	fmt.Printf("obj1 = %v\n", obj1)
	param3 := obj1.Top()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj1.GetMin()
	fmt.Printf("param_4 = %v\n", param4)
}
