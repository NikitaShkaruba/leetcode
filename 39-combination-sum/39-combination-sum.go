/*
  
  Solutions:
    1. Sort + Backtrack (Don't go into non unique): O(n*logn + (t/l)^2 * n) time, O(n + t/l) space, n is the len of candidates, l is the smallest candidate, t is the target
  
  Test cases:
  
  Testing grounds:
    [2,3,6,7], 7
     ^
    cur: [2,2,]
    possible: [
      2: [[2]],
      4: [[2,2]],
      7: [[2,2,3]]
      6: [[2,2,2]],
    ]

    
  [2,3,6,7], [], 0, 7  ------------------------------------------------------------------------------------- [], [7], 7, 7 (YES)
  ||                                                                                     ||
  [2,3,5,6], [2], 2, 7  ------------------------- [3,5,6], [2,3], 5, 7                 [3,6,7], [3], 3, 7
  || 
  [2,3,5,6], [2,2], 4, 7 -------------------------
  ||                                            ||
  [2,3,5,6], [2,2,2], 6, 7             [3,5,6], [2,2,3], 7, 7 (YES)
  ||
  [2,3,5,6], [2,2,2,2], 8, 7 (X)
  
  
  [1,2,3], 9
  ||
  ...
  ||
  [1], [1,1,1,1]
*/
func combinationSum(candidates []int, targetSum int) [][]int {
  sort.Slice(candidates, func(i, j int) bool {
    return candidates[i] < candidates[j]
  })
  
  res := make([][]int, 0)
  backtrack(candidates, nil, 0, targetSum, &res)
  
  return res
}

func backtrack(candidates, cur []int, curSum, targetSum int, res *[][]int) {
  if curSum == targetSum {
    curCopy := make([]int, len(cur))
    copy(curCopy, cur)
    *res = append(*res, curCopy)
    return
  }
  
  for i, v := range candidates {
    newSum := curSum + v
    if newSum > targetSum {
      return
    }
    
    backtrack(candidates[i:], append(cur, v), newSum, targetSum, res)
  }
}