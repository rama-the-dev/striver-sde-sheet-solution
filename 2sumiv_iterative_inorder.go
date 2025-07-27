package main

/*
653. Two Sum IV - Input is a BST
Solved
Easy
Topics
Companies
Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget4(root *TreeNode, target int) bool {
	inorder := []int{}
	stack := []*TreeNode{}
	curr := root

	// upto the point tree is not fully traversed
	// or any node is in stack keep iterating
	for curr != nil || len(stack) != 0 {
		// push all left nodes of current node to stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inorder = append(inorder, curr.Val)
		curr = curr.Right
	}

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
