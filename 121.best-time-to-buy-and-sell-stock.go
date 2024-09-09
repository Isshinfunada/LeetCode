/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
package main

import (
	"slices"
)

func maxProfit(prices []int) int {
	min := slices.Min(prices)
	max := min
	var flag bool
	for _, v := range prices {
		if v == min {
			flag = true
			continue
		}
		if flag {
			if v-min > 0 && v > max {
				max = v
			}
		}
	}
	return max - min
}

// @lc code=end
