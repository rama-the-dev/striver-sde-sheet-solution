package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget1(root *TreeNode, target int) bool {
	inorder := []int{}

	doInorder(root, &inorder)

	l, r := 0, len(inorder)-1
	currSum := 0
	for l < r {
		currSum = inorder[l] + inorder[r]
		if currSum > target {
			r -= 1
		} else if currSum < target {
			l += 1
		} else {
			return true
		}
	}
	return false
}

func doInorder(root *TreeNode, inorder *[]int) {
	if root == nil {
		return
	}
	doInorder(root.Left, inorder)
	*inorder = append(*inorder, root.Val)
	doInorder(root.Right, inorder)
}
