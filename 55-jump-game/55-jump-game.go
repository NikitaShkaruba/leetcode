/*
  
  Soultions:
    - Count max reachable index: O(n) time, O(1) space (implemented)
  
  Test cases:
    - [2,3,1,1,4]: true
    - [3,2,1,0,4]: false
    - [3,2,3,0,0,0]: true
    - [0,2]: false
    - [0]: true (!!!)
    - [543534534]: true
  
  Testing grounds:
      [0,1,2,3,4,5]
    - [3,2,3,0,0,4]
               i
    maxReachableIndex: 5
*/
func canJump(nums []int) bool {
  maxReachableIndex := 0
  
  for i := 0; i < len(nums)-1; i++ {
    if nums[i] != 0 {
      maxReachableIndex = max(maxReachableIndex, i + nums[i])
      continue
    }
    
    if maxReachableIndex <= i {
      return false
    }
  }
  
  return true
}

func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}