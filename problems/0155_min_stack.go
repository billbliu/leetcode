/*
 * @Author: bill
 * @Date: 2021-09-22 17:45:17
 * @LastEditors: bill
 * @LastEditTime: 2021-09-23 09:06:26
 * @Description: go test -v  0155_min_stack_test.go 0155_min_stack.go
 * @FilePath: /leetcode-go/problems/0155_min_stack.go
 */

package problems

type MinStack struct {
	data []int
	min  int // 最小元素位置
	top  int
}

/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack{
		data: make([]int, 10),
		min:  -1,
		top:  -1,
	}
}

func (this *MinStack) Push(val int) {
	if this.top == len(this.data)-1 {
		this.data = append(this.data, val)
		this.top++
	} else {
		this.top++
		this.data[this.top] = val
	}

	if this.min == -1 || val < this.data[this.min] {
		this.min = this.top
	}
}

func (this *MinStack) Pop() {
	if this.top != -1 {
		// 最小值是栈顶元素,重新求解最小值
		if this.top == 0 {
			this.min = -1
		} else {
			min := this.data[0]
			this.min = 0
			for i := 0; i < this.top; i++ {
				if this.data[i] < min {
					this.min = i
				}
			}
		}
		this.top--
	}
}

func (this *MinStack) Top() int {
	if this.top != -1 {
		return this.data[this.top]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.min != -1 {
		return this.data[this.min]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
