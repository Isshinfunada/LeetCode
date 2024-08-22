/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
package main

func romanToInt(s string) int {
	result := make([]int, 15)
	result = append(result, 0)

	hm := make(map[rune]int)
	hm['I'] = 1
	hm['V'] = 5
	hm['X'] = 10
	hm['L'] = 50
	hm['C'] = 100
	hm['D'] = 500
	hm['M'] = 1000

	for i, roman := range s {
		if i > 0 {
			switch roman {
			case 'I':
				result = append(result, hm['I'])
			case 'V':
				if result[i-1] == 1 {
					result = append(result, 3)
				} else {
					result = append(result, hm['V'])
				}
			case 'X':
				if result[i-1] == 1 {
					result = append(result, 8)
				} else {
					result = append(result, hm['X'])
				}
			case 'L':
				if result[i-1] == 10 {
					result = append(result, 30)
				} else {
					result = append(result, hm['L'])
				}
			case 'C':
				if result[i-1] == 10 {
					result = append(result, 80)
				} else {
					result = append(result, hm['C'])
				}
			case 'D':
				if result[i-1] == 100 {
					result = append(result, 300)
				} else {
					result = append(result, hm['D'])
				}
			case 'M':
				if result[i-1] == 100 {
					result = append(result, 800)
				} else {
					result = append(result, hm['M'])
				}
			}
		} else {
			result = append(result, hm[roman])
		}
	}

	sum := 0
	for _, num := range result {
		sum += num
	}
	return sum
}

// @lc code=end
