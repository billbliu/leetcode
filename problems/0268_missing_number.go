/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-14 11:57:25
 * @Description: go test -v 0268_missing_number_test.go 0268_missing_number.go
 * @FilePath: /leetcode-go/problems/0268_missing_number.go
 */
package problems

import "sort"

// xor
func missingNumber1(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}

// sum
func missingNumber2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i] - i
	}
	return len(nums) - sum
}

// binary search
func missingNumber3(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	mid := (left + right) / 2
	for left < right {
		if nums[mid] > mid {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return left
}
