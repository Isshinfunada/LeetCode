/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

func traversal(root *TreeNode, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		traversal(root.Right, result)
		*result = append(*result, root.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	traversal(root, &result)
	return result
}

// @lc code=end
