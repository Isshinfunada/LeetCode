/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
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

func detectCycle(head *ListNode) *ListNode {
	alreadyRunnedArray := make(map[int]*ListNode)
	currentHead := head

	for currentHead != nil {
		if _, ok := alreadyRunnedArray[currentHead.Val]; ok {
			return alreadyRunnedArray[currentHead.Val]
		}
		alreadyRunnedArray[currentHead.Val] = currentHead
		currentHead = currentHead.Next
	}
	return nil
}

// @lc code=end
