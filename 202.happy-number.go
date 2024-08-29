/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
package main

func isHappy(n int) bool {
	var values []int
	result := n
	for result != 1 {
		values = convert(result)
		result = happyLoop(values)
	}
	return true
}

// int型のスライスを渡すと、2乗した計算結果を返す
func happyLoop(value []int) int {
	r := make(map[int]int)
	for i, v := range value {
		r[i] = v * v
	}

	sum := 0
	for _, n := range r {
		sum += n
	}
	return sum
}

// 1桁ずつ分解
func convert(intnum int) []int {
	i := intnum
	var res []int
	for i != 1 {
		for i > 9 {
			res = append(res, i/10)
		}
		res = append(res, i%10)
	}

	return res
}

// @lc code=end
