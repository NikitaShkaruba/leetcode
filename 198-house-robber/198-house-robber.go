/*
  
  Solutions:
    - Dynamic programming: O(n) time, O(n) space
    - Dynamic programming + put dp array in nums: O(n) time, O(1) space
  
  Test cases:
    - [1,2,3,1]: 4
    - [2,7,9,3,1]: 12
    - [5]: 5 (!!!)
  
  Testing grounds:
    [2,7,9 ,3 ,1]
    [2,7,11,11,12]
         ^
*/
func rob(nums []int) int {
  for i := 1; i < len(nums); i++ {
    withRobbing := nums[i]
    if i-2 >= 0 {
      withRobbing += nums[i-2]
    }
    
    withoutRobbing := nums[i-1]
    
    nums[i] = max(withRobbing, withoutRobbing)
  }
  
  return nums[len(nums)-1]
}

func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}