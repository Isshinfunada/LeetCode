/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
package main

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{i, hm[num]}
		}
		hm[target-num] = i
	}
	return nil
}

// @lc code=end
