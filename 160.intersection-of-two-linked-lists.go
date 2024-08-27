/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA := headA
	currentB := headB
	alreadyRunnedArray := make(map[*ListNode]bool)
	var result *ListNode

	for currentA != nil {
		alreadyRunnedArray[currentA] = true
		currentA = currentA.Next
	}

	for currentB != nil {
		_, ok := alreadyRunnedArray[currentB]
		if ok {
			result = currentB
			return result
		}
		currentB = currentB.Next
	}
	return nil
}

// @lc code=end
