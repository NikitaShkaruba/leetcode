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
    1. DFS: O(n + m) time, O(n + m) space

  Test cases:
    - nil, [1, 2, 3]
    - [1, 2, 3], nil
    - nil, nil (!!!)

  Testing grounds:
  
    1      2           3
  /    +    \    =   /   \
 2           5      2     5
  
*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
  if root1 == nil && root2 == nil {
    return nil
  }
  
  node := &TreeNode{}
  
  var root1Left, root1Right, root2Left, root2Right *TreeNode = nil, nil, nil, nil
  if root1 != nil {
    node.Val += root1.Val
    root1Left = root1.Left
    root1Right = root1.Right
  }
  if root2 != nil {
    node.Val += root2.Val
    root2Left = root2.Left
    root2Right = root2.Right
  }
  
  node.Left = mergeTrees(root1Left, root2Left)
  node.Right = mergeTrees(root1Right, root2Right)
  
  return node
}