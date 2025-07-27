package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	stack   []*TreeNode
	forward bool
}

func NewBSTIterator(root *TreeNode, forward bool) *BSTIterator {
	it := &BSTIterator{
		stack:   []*TreeNode{},
		forward: forward,
	}
	it.pushAll(root)
	return it
}

func (it *BSTIterator) pushAll(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		if it.forward {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}

func (it *BSTIterator) hasNext() bool {
	return len(it.stack) > 0
}

func (it *BSTIterator) next() int {
	n := len(it.stack) - 1
	node := it.stack[n]
	it.stack = it.stack[:n]

	val := node.Val
	if it.forward {
		it.pushAll(node.Right)
	} else {
		it.pushAll(node.Left)
	}
	return val
}

func findTarget(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	leftIt := NewBSTIterator(root, true)
	rightIt := NewBSTIterator(root, false)

	left := leftIt.next()
	right := rightIt.next()

	for left < right {
		sum := left + right
		if sum == target {
			return true
		} else if sum < target {
			if leftIt.hasNext() {
				left = leftIt.next()
			} else {
				break
			}
		} else {
			if rightIt.hasNext() {
				right = rightIt.next()
			} else {
				break
			}
		}
	}
	return false
}
