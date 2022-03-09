/*
  
  Solutions:
    1. Init nums, backtrack: O(1) time, O(1) space
  
  Test cases:
    - 4, 1 : [] (!!!)
    - 3, 9 : [[1,2,6],[1,3,5],[2,3,4]]
  
  Testing grounds:
    [1, 2, 3, 4, 5, 6, 7, 8, 9], 7
     ^
     
    [1, 2, 3, 4, 5, 6, 7, 8, 9], [], 3, 7
    [2, 3, 4, 5, 6, 7, 8, 9], [1], 2, 6
    [3, 4, 5, 6, 7, 8, 9], [1, 2], 1, 4
    [5, 6, 7, 8, 9], [1, 2, 4], 0, 0
    
    nums, comb, k, n, res
*/
func combinationSum3(k int, n int) [][]int {
  nums := make([]int, 0, 9)
  for i := 1; i <= 9; i++ {
    nums = append(nums, i)
  }
  
  res := make([][]int, 0)
  backtrack(nums, nil, k, n, &res)
  
  return res
}

func backtrack(nums, comb []int, k, n int, res *[][]int) {
  if k == 0 && n == 0 {
    *res = append(*res, append([]int{}, comb...))
    return
  }
  
  if k == 0 || n == 0 {
    return
  }
  
  for i, v := range nums {
    newN := n-v
    if newN < 0 {
      return
    }
    
    backtrack(nums[i+1:], append(comb, v), k-1, newN, res)
  }
}