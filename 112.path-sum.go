/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	var current *TreeNode
	var diff int

	for root.Left == nil && root.Right == nil {
		current = &TreeNode{Val: root.Val, Left: root.Left, Right: root.Right}
		diff = targetSum - current.Val
	}
}

// @lc code=end
