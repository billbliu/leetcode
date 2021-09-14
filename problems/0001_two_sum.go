/*
 * @Author: bill
 * @Date: 2021-09-14 11:28:14
 * @LastEditors: bill
 * @LastEditTime: 2021-09-14 17:29:11
 * @Description: go test -v  0001_two_sum_test.go 0001_two_sum.go
 * @FilePath: /leetcode-go/problems/0001_two_sum.go
 */

package problems

// 方法一：暴力枚举
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：哈希表
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if hv, ok := hashTable[target-v]; ok {
			return []int{i, hv}
		}
		hashTable[v] = i
	}
	return nil
}
