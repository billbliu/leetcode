/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-10-08 15:27:52
 * @Description: go test -v 1206_design_skiplist.go 1206_design_skiplist_test.go
 * @FilePath: /leetcode-go/problems/1206_design_skiplist.go
 */
//  概率分层直接进行了一个随机数的对数计算，其他的跟传统跳表实现起来大同小异。
//  用的是纯跳表单元，只有值、向右、向下三个属性，没有使用数组存表，add的时候用一个prev数组记录前一个跳跃点，
//  所以不需要在结构体里用超过两个方向的形式存储，空间上比数组存表略优一点吧。

package problems

import (
	"math"
	"math/rand"
)

const (
	maxLevel = 16
	maxRand  = 65535.0
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type skipNode struct {
	value int
	right *skipNode
	down  *skipNode
}

type Skiplist struct {
	head *skipNode
}

func Constructor() Skiplist {
	left := make([]*skipNode, maxLevel)
	right := make([]*skipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &skipNode{-1, nil, nil}
		right[i] = &skipNode{20001, nil, nil}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{left[0]}
}

func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		if node.right.value > target {
			node = node.down
		} else if node.right.value < target {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	prev := make([]*skipNode, maxLevel)
	i := 0
	node := this.head
	for node != nil {
		if node.right.value >= num {
			prev[i] = node
			i++
			node = node.down
		} else {
			node = node.right
		}
	}
	n := randLevel()
	arr := make([]*skipNode, n)
	t := &skipNode{-1, nil, nil}
	for i, a := range arr {
		p := prev[maxLevel-n+i]
		a = &skipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a
	}
}

func (this *Skiplist) Erase(num int) (ans bool) {
	node := this.head
	for node != nil {
		if node.right.value > num {
			node = node.down
		} else if node.right.value < num {
			node = node.right
		} else {
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
