/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
package main

import "fmt"

func moveZeroes(nums []int) {
	cnt := 0
	for i, v := range nums {
		if v == 0 {
			fmt.Printf("cnt: %#v\n", cnt)
			cnt += 1
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			fmt.Printf("nums: %#v\n", nums)
			continue
		}
		fmt.Printf("nums: %#v\n", nums)
	}

	for i := 0; i < cnt; i++ {
		nums = append(nums, 0)
	}
}

// @lc code=end
