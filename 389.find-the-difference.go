/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
package main

func findTheDifference(s string, t string) byte {
	exists := make(map[rune]int)
	added := make(map[rune]int)
	for _, r := range s {
		cnt := exists[r]
		exists[r] = cnt + 1
	}

	for _, value := range t {
		cnt := added[value]
		added[value] = cnt + 1
	}

	for v, _ := range added {
		if exists[v] != added[v] {
			return byte(v)
		}
	}
	return 0
}

// @lc code=end
