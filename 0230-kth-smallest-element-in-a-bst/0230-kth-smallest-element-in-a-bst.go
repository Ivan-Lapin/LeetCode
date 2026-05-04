func kthSmallest(root *TreeNode, k int) int {
	var res int
	count := 0
	dfskthSmallest(root, &count, k, &res)
	return res
}

func dfskthSmallest(root *TreeNode, count *int, k int, res *int) {
	if root == nil {
		return
	}

	dfskthSmallest(root.Left, count, k, res)

	*count++
	if *count == k {
		*res = root.Val
		return
	}

	dfskthSmallest(root.Right, count, k, res)
}