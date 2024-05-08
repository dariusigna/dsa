package algo

import "github.com/dariusigna/dsa/ds"

func ArrayToTree[T comparable](arr []T) *ds.TreeNode[T] {
	if len(arr) == 0 {
		return nil
	}

	root := &ds.TreeNode[T]{Value: arr[0]}
	queue := ds.NewQueue[*ds.TreeNode[T]]()
	queue.Add(root)
	index := 1
	var zero T
	for !queue.IsEmpty() {
		node, _ := queue.Pop()
		if index >= len(arr) {
			node.Left = nil
			node.Right = nil
			break
		}

		if arr[index] != zero {
			node.Left = &ds.TreeNode[T]{Value: arr[index]}
			queue.Add(node.Left)
		}

		index++

		if arr[index] != zero {
			node.Right = &ds.TreeNode[T]{Value: arr[index]}
			queue.Add(node.Right)
		}

		index++
	}

	return root
}

func TreeToArray[T comparable](root *ds.TreeNode[T]) []T {
	if root == nil {
		return nil
	}

	var result []T
	queue := ds.NewQueue[*ds.TreeNode[T]]()
	queue.Add(root)
	var zero T
	for !queue.IsEmpty() {
		node, _ := queue.Pop()
		if node == nil {
			result = append(result, zero)
			continue
		}

		result = append(result, node.Value)
		queue.Add(node.Left)
		queue.Add(node.Right)
	}

	var i int
	for i = len(result) - 1; i >= 0; i-- {
		if result[i] != zero {
			break
		}
	}

	result = result[:i+1]
	return result
}

func TreeHeight[T comparable](root *ds.TreeNode[T]) int {
	if root == nil {
		return 0
	}

	leftHeight := TreeHeight(root.Left)
	rightHeight := TreeHeight(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func InOrderTraversal[T comparable](root *ds.TreeNode[T], f func(T)) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left, f)
	f(root.Value)
	InOrderTraversal(root.Right, f)
}

func PreOrderTraversal[T comparable](root *ds.TreeNode[T], f func(T)) {
	if root == nil {
		return
	}

	f(root.Value)
	PreOrderTraversal(root.Left, f)
	PreOrderTraversal(root.Right, f)
}

func PostOrderTraversal[T comparable](root *ds.TreeNode[T], f func(T)) {
	if root == nil {
		return
	}

	PostOrderTraversal(root.Left, f)
	PostOrderTraversal(root.Right, f)
	f(root.Value)
}

func IterativeDFSTree[T comparable](root *ds.TreeNode[T], f func(T)) {
	if root == nil {
		return
	}

	stack := ds.NewStack[*ds.TreeNode[T]]()
	stack.Push(root)
	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		f(node.Value)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

func BFSTree[T comparable](root *ds.TreeNode[T], f func(T)) {
	if root == nil {
		return
	}

	queue := ds.NewQueue[*ds.TreeNode[T]]()
	queue.Add(root)
	for !queue.IsEmpty() {
		node, _ := queue.Pop()
		f(node.Value)
		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}
}
