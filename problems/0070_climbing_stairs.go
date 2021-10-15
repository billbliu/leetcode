/*
 * @Author: bill
 * @Date: 2021-09-24 16:04:40
 * @LastEditors: bill
 * @LastEditTime: 2021-09-24 16:13:28
 * @Description: go test -v  0070_climbing_stairs.go 0070_climbing_stairs_test.go
 * @FilePath: /leetcode-go/problems/0070_climbing_stairs.go
 */
package problems

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// 动态规划（英語：Dynamic programming，简称DP）
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
