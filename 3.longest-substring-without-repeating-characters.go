/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	start := 1

	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %#v\n", i)
		hm := make(map[rune]bool)
		cnt := 1
		for ; start < len(s); start++ {
			_, ok := hm[rune(s[start])]
			fmt.Printf("start: %#v\n", start)
			if (s[i] != s[start]) && !ok {
				cnt++
				if cnt > max {
					max = cnt
				}
				hm[rune(s[start])] = true

				fmt.Printf("max: %#v\n", max)
			} else {
				if cnt > max {
					max = cnt
				}
				start = i + 1
				break
			}
		}
	}
	return max
}

// @lc code=end
