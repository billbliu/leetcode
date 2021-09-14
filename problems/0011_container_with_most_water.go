/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:14
 * @LastEditors: bill
 * @LastEditTime: 2021-09-14 17:27:48
 * @Description: go test -v  0001_two_sum_test.go 0001_two_sum.go
 * @FilePath: /leetcode-go/problems/0011_container_with_most_water.go
 */

package problems

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	sum, max := 0, 0
	for left < right {
		if height[left] < height[right] {
			sum = (right - left) * height[left]
			left++
		} else {
			sum = (right - left) * height[right]
			right--
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
