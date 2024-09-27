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
	alreadyRunnedArray := make(map[*ListNode]bool)
	currentHead := head

	for currentHead != nil {
		if _, ok := alreadyRunnedArray[currentHead]; ok {
			return currentHead
		}
		alreadyRunnedArray[currentHead] = true
		currentHead = currentHead.Next
	}
	return nil
}

// @lc code=end
