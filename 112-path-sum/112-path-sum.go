/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(node *TreeNode, targetSum int) bool {
  if node == nil {
    return false
  }
  
  if node.Left == nil && node.Right == nil {
    return node.Val == targetSum
  }
  
  if node.Left != nil {
    if hasPathSum(node.Left, targetSum - node.Val) {
      return true
    }
  }
  if node.Right != nil {
    if hasPathSum(node.Right, targetSum - node.Val) {
      return true
    }
  }
  
  return false
}