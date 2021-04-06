package array

import (
	"testing"
)

type question34 struct {
	para34
	ans34
}

// para代表参数
// one代表第一个参数
// two代表第二个参数
type para34 struct {
	one []int
	two int
}

// ans代表返回值
// one代表第一个返回值
type ans34 struct {
	one []int
}

func Test_Problem34(t *testing.T) {
	qs := []question34{
		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 8},
			ans34{[]int{3, 4}},
		},
		{
			para34{[]int{5, 7, 7, 8, 8, 10}, 6},
			ans34{[]int{-1, -1}},
		},
	}

	t.Log("————————————————————LeetCode Problen 34————————————————————")
	t.Log("————————————————————solution 1————————————————————")
	for _, q := range qs {
		t.Logf("[input]:%v		[output]:%v\n", q.para34, searchRange(q.para34.one, q.para34.two))
	}

	t.Logf("\n\n")
}
