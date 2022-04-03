/*
  
  Solutions:
    - Backtracking brute force: O(2^n) time, O(n) space
    - Tabulation: O(n^2) time, O(n) space (implemented)
  
  Test cases:
    - [1,3,5,4,7]: 2
    - [2,2,2,2,2]: 5
    - [5,4,3,2,1]: 1
    - [3]: 2
  
  Testing grounds:
    
      f(curSeq, nums)
    - [1,3,5,4,7]
                               []
       [1], [3,5,4,7]                   or               [], [3,5,4,7]   
   [1,3], [5,4,7]           or [1], [5,4,7]
   [1,3,5], [4,7]  or   [1,3], [4,7]
   [1,3,5], [7]    or   [1,3,4], [7]
   [1,3,5,7], [8,9,4]        [1,3,4,7], [8,9,3]
       
       longestSize = 1
       longestCount = 1
       
       cache key
       cache {
          7: 1
       }
       
       [1,3,5,4,7]
       [1,2,3,3,4]
       
       [2,2,2,2]
       [1,1,1,1]
       
       [5,4,3,2,1]
       [1,1,1,1,1]
       
*/
func findNumberOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    count := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
        count[i] = 1
    }
    
    longest := 1
    totalLIS := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] <= nums[j] {
              continue
            }
          
            if dp[i] < dp[j]+1 {
                dp[i] = dp[j]+1
                count[i] = count[j]
            } else if dp[i] == dp[j]+1 {
                count[i] += count[j]
            }
        }
      
        if longest < dp[i] {
            longest = dp[i]
            totalLIS = count[i]
        } else if longest == dp[i] {
            totalLIS += count[i]
        }
    }
  
    return totalLIS
}