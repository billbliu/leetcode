/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-22 14:44:39
 * @Description: go test -v  0020_valid_parentheses_test.go 0020_valid_parentheses.go
 * @FilePath: /leetcode-go/problems/0020_valid_parentheses_test.go
 */
package problems

import (
	"fmt"
	"testing"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func Test_Problem20(t *testing.T) {

	qs := []question20{

		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, q := range qs {
		_, p := q.ans20, q.para20
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValid(p.one))
	}
}
