/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
package main

func maxProfit(prices []int) int {
	dif := 0
	now := prices[0]
	for i := 1; i < len(prices); i++ {
		if now > prices[i] {
			now = prices[i]
		} else if prices[i]-now > dif {
			dif = prices[i] - now
		}
	}
	return dif
}

// @lc code=end
