/*

  Solutions:
    0. Brute force backtracking: O(a^n) time, O(a) space (We can also use memoization here)
    1. dp: O(a*n) time, O(a) space, where n is the len of coins, and a is the amount
    2. dp + only store max coin elements (If needed): O(a*n) time, O(max coin elements)
    3. sort coins + dp with preventive stop: O(n*log(n) + a*n) time, O(n + max coin elements) (implemented)
  
  Test cases:
    - coins = [2, 5], amount = 3
    - coins = [1,2,5], amount = 11
    - coins = [2], amount = 3 (!!!)
    - coins = [1], amount = 0 (!!!)
  
  Testing grounds:
    [1,2,5], 11
    
     0,1,2,3,4,5,6,7,8,9,10,11
    [0,1,1,2,2,1,2,2,3,3, 2, 3]
                             ^
                             
     [1]
     
    [2, 5], 3
    
     0, 1, 2, 3
    [0,-1, 1,-1]
              ^
                             
*/
func coinChange(coins []int, amount int) int {
  sort.Slice(coins, func(i, j int) bool {
    return coins[i] < coins[j]
  })
  
  dp := make([]int, 0, amount + 1)
  for i := 0; i < amount + 1; i++ {
    dp = append(dp, -1)
  }
  dp[0] = 0
  
  for i := 1; i <= amount; i++ {
    for _, c := range coins {
      diff := i - c
      
      if diff < 0 {
        break
      }
      
      if dp[diff] == -1 {
        continue 
      }
      
      newValue := dp[diff] + 1 
      if dp[i] != -1 {
        newValue = minInt(dp[i], newValue)
      }
      
      dp[i] = newValue
    }
  }
  
  return dp[len(dp)-1]
}

func minInt(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}