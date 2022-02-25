/*

  Solutions:
    - Brute force: O(n) time, O(n) space
  
  Test cases:
    - [1,2,3,4], m = 2, n = 2: [[1, 2], [3, 4]]
    - [1,2,3], m = 1, n = 3 : [[1,2,3]]
    - [1,2,3], m = 3, n = 1 : [[1],[2],[3]]
    - [1,2,3,4,5], m = 2, n = 2: [] (!!!)
    - [1], m = 1, n = 1: [1]
  
  Testing grounds:
    [1, 2, 3, 4]
     
           j
      [[1, 2
     i [3, 4]]]
     
     [1,1,1,1]
     
        j
     i [0,0,0,0]

*/
func construct2DArray(original []int, m int, n int) [][]int {
  if len(original) != m * n {
    return [][]int{}
  }
  
  result := make([][]int, m)
  originalIndex := 0
  
  for i := 0; i < m; i++ {
    result[i] = make([]int, n)
    
    for j := 0; j < n; j++ {
      result[i][j] = original[originalIndex]
      originalIndex++
    }
  }
  
  return result
}