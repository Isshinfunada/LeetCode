/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
package main

import (
	"fmt"
)

func isHappy(n int) bool {
	var values []int
	result := n
	for result != 1 {
		values = convert(result)
		result = happyLoop(values)
		fmt.Printf("values: %#v\n", values)
		fmt.Printf("n: %#v\n", n)
		fmt.Printf("result: %#v\n", result)
	}
	return true
}

// int型のスライスを渡すと、2乗した計算結果を返す
func happyLoop(value []int) int {
	r := make(map[int]int)
	for i, v := range value {
		r[i] = v * v
		fmt.Printf("i: %#v\n", i)
		fmt.Printf("v: %#v\n", v)
	}

	sum := 0
	for _, n := range r {
		sum += n
	}
	return sum
}

// 1桁ずつ分解
func convert(intnum int) []int {
	cnt := 0
	i := intnum

	for i >= 9 {
		i = i - 10
		cnt += 1
	}
	return []int{cnt, i}
}

// @lc code=end
