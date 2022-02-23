/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

/*
  
  Solutions:
    - Binary Search Tree traverse: O(logn) time, O(logn) space
  
  Test cases:
    - root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
    - root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
    - root = [2,1], p = 2, q = 1
  
  Testing grounds:

*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == p || root == q {
    return root
  }
  
  if root.Val > p.Val && root.Val > q.Val {
    return lowestCommonAncestor(root.Left, p, q)  
  } else if root.Val < p.Val && root.Val < q.Val {
    return lowestCommonAncestor(root.Right, p, q)  
  } else {
    return root
  }
}






