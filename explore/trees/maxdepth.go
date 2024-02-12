/*
Example 1: 104. Maximum Depth of Binary Tree

Given the root of a binary tree, find the
length of the longest path from the root to a leaf.
*/

package main

import (
	"fmt"
	"math"
)

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}

func main() {
	// Example usage:
	// Constructing a binary tree: 1 -> 2 -> 4, 1 -> 3
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	depth := maxDepth(root)
	fmt.Printf("Maximum Depth of Binary Tree: %d\n", depth)
}
