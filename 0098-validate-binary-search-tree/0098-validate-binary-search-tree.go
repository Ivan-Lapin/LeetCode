func isValidBST(root *TreeNode) bool {

	return dfcisValidBST(root, nil, nil)
}

func dfcisValidBST(root *TreeNode, min *int, max *int) bool {

    if root == nil {
		return true
	}

    if min != nil && root.Val <= *min {
        return false
    }

    if max != nil && root.Val >= *max {
        return false
    }

	return dfcisValidBST(root.Left, min, &root.Val) && dfcisValidBST(root.Right, &root.Val, max)
}
