/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

func inorderTraversal(root *TreeNode) []int {
	var result []int
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root != nil {
		helper(root.Left, result)
		*result = append(*result, root.Val)
		helper(root.Right, result)
	}
}

// @lc code=end
