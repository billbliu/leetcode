package array

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
	for {
		if left >= right {
			break
		}
		if nums[mid] > mid {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return left
}
