/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {

    if root == nil {
        return nil
    }

    result := []float64{}

    trees := []*TreeNode{root}

    for len(trees) > 0 {

        levelCount := len(trees)
        var sum float64

        for i := 0; i < levelCount; i++{

            node := trees[0]
            trees = trees[1:]
            sum += float64(node.Val)

            if node.Left != nil {
                trees = append(trees, node.Left)
            }
            if node.Right != nil {
                trees = append(trees, node.Right)
            }

        }
        result = append(result, sum/float64(levelCount))
    }

    return result
}
