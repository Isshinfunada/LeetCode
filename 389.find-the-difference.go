/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
package main

import "fmt"

func findTheDifference(s string, t string) byte {
	exists := make(map[rune]int)
	added := make(map[rune]int)
	for _, r := range s {
		cnt := exists[r]
		exists[r] = cnt + 1
		fmt.Printf("exists: %#v\n", exists[r])
	}

	for _, value := range t {
		cnt := added[value]
		added[value] = cnt + 1
		fmt.Printf("addeed: %#v\n", added[value])
	}

	for v, _ := range added {
		fmt.Printf("exists: %#v\n", exists[v])
		fmt.Printf("addeed: %#v\n", added[v])
		if exists[v] != added[v] {
			return byte(v)
		}
	}
	return 0
}

// @lc code=end
