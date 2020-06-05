package easy

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	newLeft := root.Right
	newRight := root.Left
	root.Left = newLeft
	root.Right = newRight
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
