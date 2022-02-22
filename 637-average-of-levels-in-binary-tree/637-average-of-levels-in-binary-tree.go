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
    1. Breath first search: O(n) time, O(n) space
    2. Depth first search: O(n) time, O(logn) space
  
  Test cases:
    1. [3,9,20,15,7] : [3.00000,14.50000,11.00000]
    2. [3] : [3.00000]
  
  Testing grounds:
    tree: [3,9,20,15,7]
               ^
    semi-result [{sum, amount}]: [[3, 1], [29, 2], [22, 2]]
    result: [3.0000, 14.50000, 11.0000]
*/

type SemiResultPair struct {
  sum int
  amount int
}

func averageOfLevels(root *TreeNode) []float64 {
  semiResult := helper(root, 0, []SemiResultPair{})
  
  result := make([]float64, len(semiResult))
  for i, v := range semiResult {
    result[i] = float64(v.sum) / float64(v.amount)
  }
  
  return result
}

func helper(node *TreeNode, level int, semiResult []SemiResultPair) []SemiResultPair {
  if node == nil {
    return semiResult
  }
  
  if len(semiResult) <= level {
    semiResult = append(semiResult, SemiResultPair{0, 0})
  }
  
  semiResult[level].sum += node.Val
  semiResult[level].amount++
  
  semiResult = helper(node.Left, level + 1, semiResult)
  semiResult = helper(node.Right, level + 1, semiResult)
  
  return semiResult
}

















