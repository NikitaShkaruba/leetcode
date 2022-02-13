/**
  Test cases:
    [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    [-1, -2, -3, -4]
    [1] (!!!)
    
  Sandbox:
    arr: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    max: [-1, 1, -3, 6,  2, 3, 1, -1, 4]
                     ^
                    
    arr: [5, 4, -1, 7, 8]
    max: [23,18,14,15,8]
    
    arr: [23, 18, 14, 15, 8]
          ^
    maxSum: 23
          
    [-1, -4]
    [1, -4]
    [1, 4]
    [-1, 4]
      
  Solution:
    - best possible sums O(n) time, O(1) space, where n is the len of nums
*/
func maxSubArray(nums []int) int {
  maxSum := nums[len(nums) - 1]
  
  for i := len(nums) - 2; i >= 0; i-- {
    if nums[i + 1] > 0 {
      nums[i] += nums[i + 1]
    }
    
    if nums[i] > maxSum {
      maxSum = nums[i]
    }
  }
  
  return maxSum
}