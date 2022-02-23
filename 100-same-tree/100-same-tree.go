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
    1. DFS: O(min(n, m)) time, O(n) space, where n is the len of the first tree, m is the len of the second tre
  
  Test cases:
    - [1,2,3], [1,2,3] : true
    - [1,2,1], q = [1,1,2] : false
    - [], [1, 2, 3] : false (!!!)
    - [], [] : true (!!!)
  
  Testing grounds:

*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
  if p == nil && q == nil {
    return true
  }
  if p == nil && q != nil || p != nil && q == nil {
    return false
  }
  
  if p.Val != q.Val {
    return false
  }
  
  return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}