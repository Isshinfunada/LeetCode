/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return 1 + max(minDepth(root.Left), minDepth(root.Right))
	} else {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

// @lc code=end
