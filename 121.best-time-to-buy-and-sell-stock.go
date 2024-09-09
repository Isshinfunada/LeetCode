/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
package main

import "slices"

func maxProfit(prices []int) int {
	dif := []int{0}
	var now int
	for i := len(prices) - 1; i >= 0; i-- {
		now = i
		for j := now; j >= 0; j-- {
			if prices[i]-prices[j] > 0 {
				dif = append(dif, prices[i]-prices[j])
			}
		}
	}
	return slices.Max(dif)
}

// @lc code=end
