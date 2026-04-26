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

func dfc(root *TreeNode, count int) int {
    num1, num2 := 0, 0

    if root.Right == nil && root.Left == nil {
        return count
    }
    if root.Right != nil {
        num1 = dfc(root.Right, count+1)
    }
    if root.Left != nil {
        num2 = dfc(root.Left, count+1)
    }
    if num1 >= num2 {
        return num1
    }
    return num2
}