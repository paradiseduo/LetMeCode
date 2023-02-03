/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    var xNode *TreeNode

    var getSubtreeSize func(*TreeNode) int
    getSubtreeSize = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        if node.Val == x {
            xNode = node
        }
        return 1 + getSubtreeSize(node.Left) + getSubtreeSize(node.Right)
    }

    getSubtreeSize(root)
    leftSize := getSubtreeSize(xNode.Left)
    if leftSize >= (n+1)/2 {
        return true
    }
    rightSize := getSubtreeSize(xNode.Right)
    if rightSize >= (n+1)/2 {
        return true
    }
    remain := n - leftSize - rightSize - 1
    return remain >= (n+1)/2
}
