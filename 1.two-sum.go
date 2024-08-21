/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
package main

type element struct {
	position int
	ok       bool
}

func twoSum(nums []int, target int) []int {
	var tmp element
	result := make([]int, 0)
	i := -1
	for !tmp.ok && (i < len(nums)) {
		i++
		n := target - nums[i]
		tmp = inFunctionWithMap(nums, n)
		if i == tmp.position {
			tmp.ok = false
		}
	}
	result = append(result, i, tmp.position)
	return result
}

func inFunctionWithMap(slice []int, key int) element {
	var result element
	for i, num := range slice {
		if num == key {
			result.position = i
			result.ok = true
		}
	}
	return result
}

// @lc code=end
