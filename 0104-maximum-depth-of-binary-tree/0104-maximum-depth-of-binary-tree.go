/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return dfc(root, 1)
}

func dfc(root *TreeNode, depth int) int {
    n1, n2 := depth, depth

    if root.Left != nil {
        n1 = dfc(root.Left, depth+1)
    }

    if root.Right != nil {
        n2 = dfc(root.Right, depth+1)
    }

    if n1 >= n2 {
        return n1
    }

    return n2
}
