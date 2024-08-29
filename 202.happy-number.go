/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
package main

func isHappy(n int) bool {
	record := map[int]bool{n: true}

	for n != 1 {
		var sum int
		for n != 0 {
			num := n % 10
			sum += num * num
			n /= 10
		}
		n = sum
		if record[n] {
			return false
		}
		record[n] = true
	}

	return true
}

// @lc code=end
