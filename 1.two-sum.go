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
		_, ok := hm[target-num]
		if i != 0 && ok {
			return []int{i, hm[target-num]}
		}
		hm[num] = i
	}
	return nil
}

// @lc code=end
