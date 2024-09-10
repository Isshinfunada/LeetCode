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
	// Base case
	if nil == root {
		return false
	}

	if nil == root.Left && nil == root.Right {
		// Base case
		return root.Val == targetSum

	} else {
		// General cases
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}
}

// @lc code=end
