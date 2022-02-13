/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

currentPath = "llr"
cache = {
  "l" -> 2
  "ll" -> 3
  "lr" -> 4
}

Solutions:
  - store current path + cache: O(n) time, O(n*l) space
  - depth traversal: O(n) time, O(d) space, d is the depth of the tree
*/

func isSymmetric(root *TreeNode) bool {
  if root == nil {
    return false
  }
  
  return areSymmetricNodes(root.Left, root.Right)  
}

func areSymmetricNodes(left *TreeNode, right *TreeNode) bool {
  if left == nil && right == nil {
    return true
  }
  if left == nil && right != nil {
    return false
  }
  if right == nil && left != nil {
    return false
  }
  
  if left.Val != right.Val {
    return false
  }
  
  return areSymmetricNodes(left.Left, right.Right) && areSymmetricNodes(left.Right, right.Left)
}