package array

import (
	"testing"
)

type question268 struct {
	para268
	ans268
}

// para代表参数
// one代表第一个参数
type para268 struct {
	one []int
}

// ans代表返回值
// one代表第一个返回值
type ans268 struct {
	one int
}

func Test_Problem268(t *testing.T) {
	qs := []question268{
		{
			para268{[]int{3, 0, 1}},
			ans268{2},
		},
		{
			para268{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
			ans268{8},
		},
	}

	t.Log("————————————————————LeetCode Problen 268————————————————————")
	t.Log("————————————————————solution 1————————————————————")
	for _, q := range qs {
		t.Logf("[input]:%v		[output]:%v\n", q.para268, missingNumber1(q.para268.one))
	}

	t.Log("————————————————————solution 2————————————————————")
	for _, q := range qs {
		t.Logf("[input]:%v		[output]:%v\n", q.para268, missingNumber2(q.para268.one))
	}

	t.Log("————————————————————solution 3————————————————————")
	for _, q := range qs {
		t.Logf("[input]:%v		[output]:%v\n", q.para268, missingNumber3(q.para268.one))
	}
	t.Logf("\n\n")
}
