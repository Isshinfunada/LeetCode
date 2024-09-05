/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current1 := l1
	current2 := l2
	var result *ListNode
	up := 0

	for current1 != nil || current2 != nil {
		if current1 == nil {
			current1 = &ListNode{Val: 0, Next: current1}
		}
		if current2 == nil {
			current2 = &ListNode{Val: 0, Next: current2}
		}
		fmt.Printf("current1Val: %#v\n", current1.Val)
		fmt.Printf("current2Val: %#v\n", current2.Val)
		fmt.Printf("up: %#v\n", up)
		result = &ListNode{Val: (current1.Val + current2.Val + up) % 10, Next: result}
		if (current1.Val + current2.Val + up) >= 10 {
			up = 1
			if current1.Next == nil || current2.Next == nil {
				result.Next.Val = up
			}
		} else {
			up = 0
		}
		current1 = current1.Next
		current2 = current2.Next
		fmt.Printf("result: %#v\n", result)
	}
	return result
}

// @lc code=end
