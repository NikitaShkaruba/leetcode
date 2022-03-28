/*
    
    Solutions:
        - Backtracking: O(n^n) time, O(n^n) space
        - Sort + Backtracking: O(n^n) time, O(m^n) space, but faster then first one
        - Dynamic programming: O(t*n) time, O(t) space, where t is the target (implemented)
    
    Test cases:
        - nums = [1,2,3], target = 4 : 7
        - nums = [0], target = 3 : 0
    
    Testing grounds:
        - [3,1,2], 0, 7
          /                \                       \
     [3, 1, 2], 3, 7     [1,2], 1, 7              [2], 2, 7    
         |
     [3,1,2], 6, 7
         |             \
     [3,1,2],9,7(X)    [1,2],7,7  
        
        
        f(nums, curSum, target)
        
        [0.1,2,3,4]
        [1,1,2,4,7]

*/
func combinationSum4(nums []int, target int) int {
  dp := make([]int, target+1)
  dp[0] = 1
  
  for i := 1; i <= target; i++ {
    for _, v := range nums {
      j := i - v
      if j >= 0 {
        dp[i] += dp[j]
      }
    }
  }
  
  return dp[target]
}