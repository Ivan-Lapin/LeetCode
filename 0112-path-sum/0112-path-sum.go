/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    
    if root == nil {
        return false
    }

    return wrapperhasPathSum(root, targetSum)
    
}

func wrapperhasPathSum(root *TreeNode, targetSum int) bool {
    targetSum -= root.Val
    if root.Right == nil && root.Left == nil{
        if targetSum == 0{
            return true
        } else {
            return false
        }
    }

    var b1, b2 bool

    if root.Left != nil {
        b1 = wrapperhasPathSum(root.Left, targetSum)
    }

    if root.Right != nil {
        b2 = wrapperhasPathSum(root.Right, targetSum)
    }

    return b1 || b2
}
