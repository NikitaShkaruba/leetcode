/*
  
  Solutions:
    - Backtracking: O(n^k) time, O(k) space, without the resulted array space
  
  Test cases:
    - 4, 2: [[2,4],[3,4],[2,3],[1,2],[1,3],[1,4]]
    - 5, 1: [[1],[2],[3],[4],[5]]
    - 2, 2: [[1,2]]
    - 1, 1: [[1]]
  
  Testing grounds:
    [1,2,3,4]
     f s
     
    [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]

                    nums=[1,2,3,4], cur=[], need=2
    nums=[2,3,4], cur=[1], need=2  -- nums=[1,3,4], cur=[2], need=2 -- nums=[1,2,4], cur=[3], need=2 -- nums=[1,2,3], cur=[4], need=2
                ||
    nums=[3,4], cur=[1,2], need=2 -- nums=[2,4], cur=[1,3], need=2 -- nums=[2,3], cur=[1,4], need=2
    
    [1], cur=[], need=1
          ||
    [], cur=[1], need=1
    
    n = 4, k = 2
    nums=[1,2,3,4], cur=[], k=2
            ^
    nums=[2,3,4], cur=[1], k=2 -- nums=[2,1,3,4], cur=[2], k=2
    
    nums=[2,3,4], cur=[1], k=2
    
    nums=[4,3,1], cur=[2], k=2
    
    1234
    - 2134
    - 3124
    - 4321
    
    2314
    
*/
func combine(n int, k int) [][]int {
  nums := make([]int, 0, n)
  for i := 1; i <= n; i++ {
    nums = append(nums, i)
  }
  
  comb := make([]int, 0, k)
  res := [][]int{}
  backtrack(nums, comb, k, &res)
  
  return res
}

func backtrack(nums, comb []int, k int, res *[][]int) {
  if len(comb) == k {
    combCopy := make([]int, len(comb))
    copy(combCopy, comb)
    *res = append(*res, combCopy)
    return
  }
  
  for i := range nums {
    if len(comb) > 0 && nums[i] < comb[len(comb)-1] {
      continue
    }
    
    nums[0], nums[i] = nums[i], nums[0]
    backtrack(nums[1:], append(comb, nums[0]), k, res)
    nums[0], nums[i] = nums[i], nums[0]
  }
}