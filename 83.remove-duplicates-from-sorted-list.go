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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := &ListNode{Val: head.Val}
	current := result

	for head.Next != nil {
		head = head.Next
		if head.Val != current.Val {
			current.Next = &ListNode{Val: head.Val}
			current = current.Next
		}
	}
	return result
}

// @lc code=end
