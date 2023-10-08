/*
104. Maximum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along
the longest path from the root node down to the farthest leaf node.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


//Binary Tree
*/

package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int       // Value of the node
	Left  *TreeNode // Pointer to the left child node
	Right *TreeNode // Pointer to the right child node
}

// maxDepth calculates the maximum depth of the binary tree rooted at 'root'
func maxDepth(root *TreeNode) int {
	// Base case: if 'root' is nil (i.e., we've reached a leaf node or an empty tree),
	// return 0 to indicate that there are no more levels below this point.
	if root == nil {
		return 0
	}

	// Recursively calculate the maximum depth of the left subtree.
	leftDepth := maxDepth(root.Left)

	// Recursively calculate the maximum depth of the right subtree.
	rightDepth := maxDepth(root.Right)

	// The maximum depth of the tree rooted at 'root' is the maximum of the depths of
	// its left and right subtrees, plus one (for the current level).
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	// Example 1
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	// Calculate the maximum depth of the tree rooted at 'root1'
	result1 := maxDepth(root1)
	fmt.Println(result1) // Output: 3
	// Example 2
	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}

	// Calculate the maximum depth of the tree rooted at 'root2'
	result2 := maxDepth(root2)
	fmt.Println(result2) // Output: 2
}
