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
    1. DFS: O(n) time, O(logn) space
  
  Test cases:
    1. [3,9,20,null,null,15,7] : 2
    2. [2,null,3,null,4,null,5,null,6] : 5
    3. nil => 0 (!!!)
    4. [1] : 1 (!!!)
  
  Testing grounds:

*/

func minDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }
  
  return dfs(root, 1)
}

func dfs(node *TreeNode, level int) int {
  if node.Left == nil && node.Right == nil {
    return level
  }
  
  leftMinDepth := math.MaxInt32
  if node.Left != nil {
    leftMinDepth = dfs(node.Left, level + 1)
  }
  
  rightMinDepth := math.MaxInt32
  if node.Right != nil {
    rightMinDepth = dfs(node.Right, level + 1)
  }
  
  if leftMinDepth < rightMinDepth {
    return leftMinDepth
  } else {
    return rightMinDepth
  }
}