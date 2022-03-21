/*

  Solutions:
    1. Brute force: O(n*2^n) time, O(n) space
    2. Brute force + cache: O(n + 2^n) time, O(n) space
    3. DP: O(n^2) time, O(n) space

  Test cases:
    - [0,1,2,3,4,5] : 6
    - [0,1,0,3,2,3] : 4
    - [7,7,7,7,7,7,7] : 1
    - [3,2,1] : 1
  
  Testing grounds:
       [0,3,1,6,2,2,7]
       /             \
[0,3,1,6,2,2,7]      [3,1,6,2,2,7]
   ^                  ^
   
  [10,9,2,5,3,7,101,18]
  [1 ,1,1,2,2,3,4  ,4 ]
                    ^
*/
func lengthOfLIS(nums []int) int {
  dp := make([]int, len(nums))
  
  for i := 0; i < len(nums); i++ {
    dp[i] = 1
  }
  
  maxlen := 1
  
  for i := 1; i < len(nums); i++ {
    for j := 0; j < i; j++ {
      if nums[i] > nums[j] {
        dp[i] = max(dp[i], dp[j] + 1)
      }
    }
    
    maxlen = max(maxlen, dp[i])
  }
    
  return maxlen
}

func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}