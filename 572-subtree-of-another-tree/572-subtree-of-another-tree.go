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
    1. Brute force: O(n*m) time, O(n) space
  
  Test cases:
    - [1, 2, 3, 4], [3, 4]: true
    - [1, 2, 3], [1, 2, 3]: true (!!!)
    - [3,4,5,1,2,0], [4,1,2]: false
    - [1], [1]: true
  
  Testing grounds:

*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil {
    return false
  }
  
  if root.Val == subRoot.Val {
    if (isSubtreeHelper(root, subRoot)) {
      return true
    }
  }
  
  return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSubtreeHelper(rootNode, subRootNode *TreeNode) bool {
  if rootNode == nil && subRootNode == nil {
    return true
  }
  if rootNode == nil && subRootNode != nil || rootNode != nil && subRootNode == nil {
    return false
  }
  
  if rootNode.Val != subRootNode.Val {
    return false
  }
  
  return isSubtreeHelper(rootNode.Left, subRootNode.Left) && isSubtreeHelper(rootNode.Right, subRootNode.Right)
}