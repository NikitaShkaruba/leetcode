/*

  Solutions:
    - Recursion + cache (Top-down): O(m*n) time, O(m*n) space
    - Tabulation (Bottom-up): O(m*n) time, O(m*n) space (implemented)
    - Tabulation + only keep top row (Bottom-up): O(m*n) time, O(n) space
  
  Test cases:
    - 3, 7 : 28
    - 3, 2: 3 (testing)
    - 7, 1: 1
    - 1, 1: 1 (!!!)
  
  Testing grounds:
    @x [1,1]
    xx [1,2]
    x@ [1,3]
    
    @xx [1,1,1]
    xxx [1,2,3]
    xxx [1,3,6]
    xxx [1,4,10]
    xxx [1,5,15]
    xxx [1,6,21]
    xx@ [1,7,28]
*/
func uniquePaths(m int, n int) int {
  dp := make([][]int, m)
  for i := 0; i < m; i++ {
    dp[i] = make([]int, n)
  }
  dp[0][0] = 1
  
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      paths := 0   
      if i - 1 >= 0 {
        paths += dp[i-1][j]
      }
      if j - 1 >= 0 {
        paths += dp[i][j-1]
      }
      
      dp[i][j] += paths
    }
  }
  
  return dp[m-1][n-1]
}