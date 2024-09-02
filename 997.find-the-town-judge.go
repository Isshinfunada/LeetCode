/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
package main

func findJudge(n int, trust [][]int) int {
	judgement := make(map[int]int)
	trustflag := make(map[int]bool)
	var cnt int

	for _, v := range trust {
		cnt = judgement[v[1]]
		judgement[v[1]] = cnt + 1
		trustflag[v[0]] = true
	}

	for i := 1; i <= n; i++ {
		if judgement[i] != n-1 {
			continue
		}
		trustcheck := trustflag[i]
		if !trustcheck {
			return i
		}
	}
	return -1
}

// @lc code=end
