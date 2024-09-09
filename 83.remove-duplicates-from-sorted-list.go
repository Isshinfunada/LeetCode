/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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

func deleteDuplicates(head *ListNode) *ListNode {
	hm := make(map[int]bool)
	var result *ListNode
	for head.Next != nil {
		if _, ok := hm[head.Val]; !ok {
			result = &ListNode{Val: head.Val, Next: head.Next}
			result = result.Next
			fmt.Printf("result: %#v\n", result)
		}
		hm[head.Val] = true
		fmt.Printf("hm: %#v\n", hm)
		head = head.Next
		fmt.Printf("head: %#v\n", head)
	}
	return result
}

// @lc code=end
