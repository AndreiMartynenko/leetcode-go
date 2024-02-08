/*
BINARY TREES
BFS

199. Binary Tree Right Side View

Given the root of a binary tree, imagine yourself standing
on the right side of it, return the values of the nodes you
can see ordered from top to bottom.

Example 1:

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:

Input: root = [1,null,3]
Output: [1,3]
Example 3:

Input: root = []
Output: []

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView returns the right side view of a binary tree.
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		// Traverse each level and add the rightmost node to the result.
		for i := 0; i < size; i++ {
			node := queue[i]

			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Remove processed nodes from the queue.
		queue = queue[size:]
	}

	return result
}

/*
func main() {
	// Example usage:
	// Constructing a binary tree: 1 -> 2 -> 5, 1 -> 3 -> nil, 2 -> 4 -> nil
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Get the right side view of the binary tree.
	result := rightSideView(root)

	// Print the result.
	fmt.Printf("Right Side View of Binary Tree: %v\n", result)
}
*/
