/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := calcDepth(root.Left) - calcDepth(root.Right)
	if diff <= 1 && diff >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func calcDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(calcDepth(root.Left), calcDepth(root.Right))
}

// @lc code=end
