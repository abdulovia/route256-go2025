package main

import (
	"fmt"
	"os"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree() *TreeNode {
	var n int
	fmt.Fscan(os.Stdin, &n) // Количество узлов в дереве

	if n == 0 {
		return nil
	}

	nodes := make([]*TreeNode, n+1) // Массив для хранения узлов
	for i := 1; i <= n; i++ {
		nodes[i] = &TreeNode{}
	}

	for i := 1; i <= n; i++ {
		var val, left, right int
		fmt.Fscan(os.Stdin, &val, &left, &right)
		nodes[i].Val = val
		if left != -1 {
			nodes[i].Left = nodes[left]
		}
		if right != -1 {
			nodes[i].Right = nodes[right]
		}
	}

	return nodes[1] // Корень дерева
}

func main() {
	fmt.Println("Введите количество узлов, затем данные узлов в формате (значение, левый_ребенок, правый_ребенок):")
	fmt.Println("(Используйте -1 для отсутствующих детей)")
	root := buildTree()

	fmt.Println("Корень дерева:", root.Val)
}
