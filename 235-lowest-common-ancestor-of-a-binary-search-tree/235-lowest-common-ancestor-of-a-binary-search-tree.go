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
    - DFS with remembering path: O(n) time, O(n) space
  
  Test cases:
    - root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
    - root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
    - root = [2,1], p = 2, q = 1
  
  Testing grounds:

*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  pPath, qPath := findPath(root, p, q, []*TreeNode{}, nil, nil)
  
  for len(pPath) != len(qPath) {
    if len(pPath) > len(qPath) {
      pPath = pPath[:len(pPath)-1]
    } else {
      qPath = qPath[:len(qPath)-1]
    }  
  }
  
  for pPath[len(pPath)-1] != qPath[len(qPath)-1] {
    pPath = pPath[:len(pPath)-1]
    qPath = qPath[:len(qPath)-1]
  }
  
  return pPath[len(pPath)-1]
}

func findPath(node, p, q *TreeNode, path, pPath, qPath []*TreeNode) ([]*TreeNode, []*TreeNode) {
  if node == nil {
    return pPath, qPath
  }
  if pPath != nil && qPath != nil {
    return pPath, qPath
  }
  
  path = append(path, node)
  
  if node == p && pPath == nil {
    pPath = copyPathSlice(path)
  } else if node == q && qPath == nil {
    qPath = copyPathSlice(path)
  }
  
  pPath, qPath = findPath(node.Left, p, q, path, pPath, qPath)
  pPath, qPath = findPath(node.Right, p, q, path, pPath, qPath)
  
  return pPath, qPath
}

func copyPathSlice(path []*TreeNode) []*TreeNode {
  result := make([]*TreeNode, 0, len(path))
  
  for _, p := range path {
    result = append(result, p)
  }
  
  return result
} 








