/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
package main

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, n := range nums {
		_, exist := count[n]
		if exist {
			count[n] += 1
		} else {
			count[n] = 1
		}
	}

	most := 0
	var result int
	for i, m := range count {
		if m > most {
			result = i
			most = m
		}
	}
	return result
}

// @lc code=end
