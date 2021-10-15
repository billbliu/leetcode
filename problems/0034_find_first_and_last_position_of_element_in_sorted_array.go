/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:30
 * @LastEditors: bill
 * @LastEditTime: 2021-09-27 10:40:04
 * @Description: go test -v  0034_find_first_and_last_position_of_element_in_sorted_array_test.go 0034_find_first_and_last_position_of_element_in_sorted_array.go
 * @FilePath: /leetcode-go/problems/0034_find_first_and_last_position_of_element_in_sorted_array.go
 */
package problems

func searchRange(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

func searchRange1(nums []int, target int) []int {
	return []int{searchFirstGreaterOrEqualElement(nums, target), searchLastLessOrEqualElement(nums, target)}
}

// mid=(low+high)/2 这种写法是有问题的。因为如果 low 和 high 比较大的话，两者之和就有可能会溢出。
// 改进的方法是将 mid 的计算方式写成 low+(high-low)/2。更进一步，如果要将性能优化到极致的话，
// 我们可以将这里的除以 2 操作转化成位运算 low+ ((high-low)>>1)。因为相比除法运算来说，计算机处理位运算要快得多。

// 二分查找第一个等于target的元素，时间复杂度为O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	index := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 找到第一个与target相等的元素
			if mid == 0 || nums[mid-1] != target {
				index = mid
				break
			} else {
				right = mid - 1
			}
		}
	}
	return index
}

// 二分查找最后一个等于target的元素，时间复杂度为O(logn)
func searchLastEqualElement(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	index := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 找到最后一个与target相等的元素
			if mid == len(nums)-1 || nums[mid+1] != target {
				index = mid
				break
			} else {
				left = mid + 1
			}
		}
	}
	return index
}

// 二分查找第一个大于等于target的元素，时间复杂度为O(logn)
func searchFirstGreaterOrEqualElement(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	index := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				index = mid
				break
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

// 二分查找最后一个小于等于target元素，时间复杂度为O(logn)
func searchLastLessOrEqualElement(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	index := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				index = mid
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
