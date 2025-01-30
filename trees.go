package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsTree(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	left := dfsTree(node.Left)
	right := dfsTree(node.Right)
	return append(append(left, node.Val), right...)
}

func bfsTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		result = append(result, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	result := dfsTree(root)
	fmt.Printf("DFS traversal of the binary tree: %v\n", result)
}

