/*

  Solutions:
    - Backtrack: O(2^n) time, O(n) space
    - Dynamic programming: O(2^n) time, O(2^n) space
  
  Test cases:
    - [5,3], 1 : 0
    - [5,3], -1 : 0
    - [0,0], 2: 0
    - [1,2,3], 6: 1
    
  Testing grounds:
    [1, 2, 3, 3, 3, 4, 5], 15
    
    [1, 2, 3, 4, 5], "", 15
    [2, 3, 4, 5], "+1", 16                             -- [2, 3, 4, 5], "-1", 14
    [3, 4, 5], "+1+2", 18 -- [3, 4, 5], "+1-2", 14
    [], "+1+2+3-4-5", 0
    
    [1, 2, 3, 3, 3, 3, 3, 3, 3, 2, 1], 4
    
    {
      3: ["+1+2", "3", "3", "3"]
    }
    
    
    [1, 2, 3, 3, 3, 4, 5], 15
    {
      10: n,
      20: m,
      ...
    }
    
    [1, 2, 3, 3, 3, 4]
    
    [1, 2, 3, 3, 3]
    
    [1, 2, 3, 3]
    
    [1, 2, 3]
    {
      6: 1,
      4: 1,
      2: 1,
      1: 2,
      0: 4,
      -1: 3,
      -2: 1,
      -4: 1,
      -6: 1
    }
    
    [1, 2]
    {
      3: 1,
      1: 2,
      -1: 2,
      -3: 1
    }
    
    [1]
    {
      1: 1,
      -1: 1,
    }
    
    [1, 2, 3, 3, 3, 4, 5], 15
     ^
*/
func findTargetSumWays(nums []int, target int) int {
  if len(nums) == 0 {
    if target == 0 {
      return 1
    } else {
      return 0
    }
  }
  
  newNums := nums[1:]
  return findTargetSumWays(newNums, target + nums[0]) + findTargetSumWays(newNums, target - nums[0])
}








