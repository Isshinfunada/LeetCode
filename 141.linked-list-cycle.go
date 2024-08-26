/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
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

func hasCycle(head *ListNode) bool {
	alreadyRunnedArray := make(map[*ListNode]bool)
	currentHead := head
	for currentHead != nil {
		if alreadyRunnedArray[currentHead] {
			return true
		}
		alreadyRunnedArray[currentHead] = true
		currentHead = currentHead.Next

		fmt.Printf("alreadyRunnedArray: #%v\n", alreadyRunnedArray)
		fmt.Printf("currentHead: #%v\n", currentHead)
	}
	return false
}

// @lc code=end
