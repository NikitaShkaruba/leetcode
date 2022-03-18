/*
  
  Solutions:
    - Brute force: O(n^2) time, O(1) space
    - Combo: O(n) time, O(1) space
  
  Test cases:
    - [2,3,-2,4] : 6
    - [10, -7, -7]: 49
    - [-2,0,-1] : 0
    - [-10] : -10
  
  Testing grounds:
    nums       = [2,3,0,-2,-6,-3,-5,4]
    max_so_far = 6
    min_so_far = 2
    max = nil

    [2, 3,-2,4]
    [3,  ,  , ]
    [-2, ,  , ]
    [4,  ,  , ]
    
*/
func maxProduct(nums []int) int {
  maxProduct := nums[0]
  
  maxSoFar := nums[0]
  minSoFar := nums[0]
  for i := 1; i < len(nums); i++ {
    newMaxSoFar := max(nums[i], max(maxSoFar * nums[i], minSoFar * nums[i]))
    minSoFar = min(nums[i], min(maxSoFar * nums[i], minSoFar * nums[i]))
    maxSoFar = newMaxSoFar
    
    if maxSoFar > maxProduct {
      maxProduct = maxSoFar
    }
  }
  
  return maxProduct
}

func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}