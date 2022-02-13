/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
  
    l := inorderTraversal(root.Left)
    l = append(l, root.Val)
  
    r := inorderTraversal(root.Right)
    l = append(l, r...)
  
    return l
}