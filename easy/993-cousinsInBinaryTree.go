package easy

// TreeNode definition
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var depth1, depth2, parent1, parent2 int
	dfs(root, x, 0, -1, &parent1, &depth1)
	dfs(root, y, 0, -1, &parent2, &depth2)
	return depth1 > 1 && depth1 == depth2 && parent1 != parent2
}

func dfs(root *TreeNode, val, depth, last int, parent, res *int) {
	if root == nil {
		return
	}
	if root.Val == val {
		*res = depth
		*parent = last
		return
	}
	depth++
	dfs(root.Left, val, depth, root.Val, parent, res)
	dfs(root.Right, val, depth, root.Val, parent, res)
}
