package ds

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
