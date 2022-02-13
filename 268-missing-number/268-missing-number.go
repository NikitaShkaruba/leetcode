/*
  
  Test cases:
    - [3, 0, 1]: 2
    - [0, 1]: 2
    - [9,6,4,2,3,5,7,0,1]: 8
    - [1]: 0 (!!!)
  
  Solutions:
    - Iterate and shit: O(n) time, O(1) space
  
  Testing ground: 
    - [-3, -5, 1, -4]:
  
    - [-3, -5, 1, -4]: 2
               ^
    
    - [-3, -1, 2]: 2
               ^
               
    - [-9,-6,-4,-2,-3,-5,-7,-11,1,-10]: 8
                                ^
                                
    - [1, -2]: 0 (!!!)
       ^
*/
func missingNumber(nums []int) int {
  nums = append(nums, len(nums) + 1)
  
  zeroSubstituteNum := len(nums) + 1
  
  for i := 0; i < len(nums) - 1; i++ {
    v := int(math.Abs(float64(nums[i])))
    
    if v == zeroSubstituteNum {
      nums[0] *= -1
    } else if nums[v] == 0 {
      nums[v] = -zeroSubstituteNum
    } else {
      nums[v] *= -1
    }
  }
  
  for i := range nums {
    if nums[i] >= 0 {
      return i
    }
  }
  
  // Shouldn't happen
  return -1
}