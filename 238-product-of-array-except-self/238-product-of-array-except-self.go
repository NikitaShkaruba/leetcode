/*
  
  Solutions:
    1. Left products and right products: O(n) time, O(n) space
    2. Left products and right products + stored in the result: O(n) time, O(1) space
  
  Test cases:
    - [1,2] : [2, 1]
    - [1,2,3,4] : [24, 12, 8, 6]
    - [-1,1,0,-3,3] : [0,0,9,0,0]
  
  Testing grounds:
    nums:   [2,2,3,4]
    left:   [1,2,4,12]
    right:  [24,12,4,1]
    answer: [24,24,16,12]
                   ^
                   
    nums:   [2,2,3,4]
    answer: [1,2,4,12] 
    mul: 12
    
    nums: [-1,1,0,-3,3]
    answer: [0,0,9,0,0]
    mul: 1
*/
func productExceptSelf(nums []int) []int {
  answer := make([]int, len(nums))
  
  mul := 1
  for i := 0; i < len(nums); i++ {
    answer[i] = mul
    mul *= nums[i]
  }
  
  mul = 1
  for i := len(nums) - 1; i >= 0; i-- {
    answer[i] *= mul
    mul *= nums[i]
  }
  
  return answer
}