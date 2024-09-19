/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
package main

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	for i, n := range nums {
		if target == n {
			return i
		} else if i < len(nums)-1 && target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}
	return len(nums)
}

// @lc code=end
