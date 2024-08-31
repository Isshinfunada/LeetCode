/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
package main

func isAnagram(s string, t string) bool {
	original := make(map[rune]int)
	anagram := make(map[rune]int)
	result := true

	for _, v := range s {
		tmp := original[v]
		original[v] = tmp + 1
	}

	for _, r := range t {
		tmp := anagram[r]
		anagram[r] = tmp + 1
	}

	for val, _ := range original {
		if len(original) != len(anagram) {
			result = false
		}
		if anagram[val] != original[val] {
			result = false
		}
	}
	return result
}

// @lc code=end
