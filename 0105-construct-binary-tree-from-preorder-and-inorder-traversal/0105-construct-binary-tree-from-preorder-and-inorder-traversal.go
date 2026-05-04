func buildTree(preorder []int, inorder []int) *TreeNode {

	inMap := make(map[int]int)

	for i, v := range inorder {
		inMap[v] = i
	}

	preIndex := 0

	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preIndex]
        preIndex++

		root := &TreeNode{Val: rootVal}

		mid := inMap[rootVal]

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)

		return root
	}

	return dfs(0, len(inorder)-1)
}