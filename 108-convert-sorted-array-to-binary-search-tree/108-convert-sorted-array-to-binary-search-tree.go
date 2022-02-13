/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 
 
 Test cases:
  * [-10,-3,0,5,9]
  * [-10, -5, -3, 0, 5, 7, 9]
  * [1] (!!!)
  
  
  Testing grounds:
    [-10, -5, -3, 0, 5, 7, 9]
  [-10, -5, -3], [0], [5, 7, 9]
          0
      -5     7
  -10  -3  5   9
  
  Solutions:
    * Brute force: O(n) time, O(n) space
 */
func sortedArrayToBST(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  }
  
  rootIndex := int(math.Floor(float64(len(nums) / 2)))
  
  var left *TreeNode
  if rootIndex > 0 {
    left = sortedArrayToBST(nums[:rootIndex])
  }
  
  var right *TreeNode
  if rootIndex + 1 < len(nums) {
    right = sortedArrayToBST(nums[rootIndex+1:]) 
  }
  
  return &TreeNode{
    Val: nums[rootIndex],
    Left: left,
    Right: right,
  }
}