func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	m := make(map[int][]int)
	BinTreeLevel(root, 0, m)

	lenght := len(m)

	res := make([][]int, lenght)
	for i := 0; i < lenght; i++ {
		res[i] = m[i]
	}

	return res
}

func BinTreeLevel(root *TreeNode, level int, m map[int][]int) {
    if root == nil {
        return 
    }
	m[level] = append(m[level], root.Val)

    if root.Left != nil {
        BinTreeLevel(root.Left, level+1, m)
    }
    if root.Right != nil {
        BinTreeLevel(root.Right, level+1, m)
    }
}