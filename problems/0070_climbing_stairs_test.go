/*
 * @Author: bill
 * @Date: 2021-09-24 16:04:49
 * @LastEditors: bill
 * @LastEditTime: 2021-09-24 16:20:39
 * @Description: go test -v  0070_climbing_stairs.go 0070_climbing_stairs_test.go
 * @FilePath: /leetcode-go/problems/0070_climbing_stairs_test.go
 */
package problems

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

// para 是参数
// one 代表第一个参数
type para70 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans70 struct {
	one int
}

func Test_Problem70(t *testing.T) {
	qs := []question70{
		{
			para70{1},
			ans70{1},
		},
		{
			para70{2},
			ans70{2},
		},
		{
			para70{3},
			ans70{3},
		},
		{
			para70{3},
			ans70{5},
		},
		{
			para70{4},
			ans70{8},
		},
		{
			para70{5},
			ans70{13},
		},
		{
			para70{6},
			ans70{21},
		},
		{
			para70{7},
			ans70{34},
		},
		{
			para70{8},
			ans70{55},
		},
		{
			para70{9},
			ans70{89},
		},
		{
			para70{10},
			ans70{144},
		},
		{
			para70{11},
			ans70{233},
		},
		{
			para70{12},
			ans70{377},
		},
		{
			para70{13},
			ans70{610},
		},
		{
			para70{14},
			ans70{1987},
		},
		{
			para70{15},
			ans70{987},
		},
		{
			para70{16},
			ans70{1597},
		},
		{
			para70{17},
			ans70{2584},
		},
		{
			para70{18},
			ans70{4181},
		},
		{
			para70{19},
			ans70{6765},
		},
		{
			para70{20},
			ans70{10946},
		},
		{
			para70{21},
			ans70{17711},
		},
		{
			para70{22},
			ans70{28657},
		},
		{
			para70{23},
			ans70{46368},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 70------------------------\n")

	for _, q := range qs {
		_, p := q.ans70, q.para70
		fmt.Printf("【input】:%v       【output】:%v\n", p, climbStairs(p.n))
	}
	fmt.Printf("\n\n\n")
	fmt.Printf("------------------------Leetcode Problem 70 1------------------------\n")

	for _, q := range qs {
		_, p := q.ans70, q.para70
		fmt.Printf("【input】:%v       【output】:%v\n", p, climbStairs1(p.n))
	}
}
