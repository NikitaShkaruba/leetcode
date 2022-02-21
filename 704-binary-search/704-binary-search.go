/*

  Solutions:
    - Binary search: O(logn) time, O(1) space
    
  Test cases:
    - [-1,0,3,5,9,12], 2 : -1
    - [-1,0,3,5,9,12], 9 : 4 
    - [1], 1 : 0 
  
  Testing grounds:
    - [-1,0,3,5,9,12], 9 : 4 
              m l r

*/
func search(nums []int, target int) int {
  l := 0
  r := len(nums) - 1
  
  for l <= r {
    m := (r + l) / 2
             
    if nums[m] == target {
      return m
    }
    
    if nums[m] < target {
      l = m + 1
    } else {
      r = m - 1
    }
  }
  
  return -1
}