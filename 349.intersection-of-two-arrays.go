/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
package main

import "slices"

func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return nil
	}

	var result []int
	if len(nums1) > len(nums2) {
		for _, n := range nums2 {
			if !slices.Contains(result, n) && slices.Contains(nums1, n) {
				result = append(result, n)
			}
		}
	}

	if len(nums2) > len(nums1) {
		for _, n := range nums1 {
			if !slices.Contains(result, n) && slices.Contains(nums2, n) {
				result = append(result, n)
			}
		}
	}
	return result
}

// @lc code=end
