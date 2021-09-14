/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-14 11:49:58
 * @Description: go test -v  0001_two_sum_test.go 0001_two_sum.go
 * @FilePath: /leetcode-go/problems/0001_two_sum_test.go
 */
package problems

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para代表参数
// one代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans代表返回值
// one代表第一个返回值
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum1(p.nums, p.target))
	}

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum2(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
