/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

  Solutions:
    1. DFS: O(n) time, O(n) space
  
  Test cases:
    - [1,2,3]: [1,3,2]
    - [1]: [1] (!!!)
    - [] (!!!)

*/
func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }
  
  root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
  
  return root
}