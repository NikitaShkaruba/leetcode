/*

  Solutions:
    1. HashMap + backtracking + remove duplicates: O(a lot) time, O(a lot) space
    2. HashMap + backtracking: O(n + n^(t/l)) time, O(n + t/l + n^(t/l)) space
    3. Sort + Backtracking (Coz everything is sorted, we can easily break out of iterations): O(n*logn + n^(t/l)) time, O(n + t/l + n^(t/l))
  
  Test cases:
    - [10,1,2,7,6,1,5] : [[1,1,6],[1,2,5],[1,7],[2,6]]
    - [1], 5 : [] (!!!)
    - [30], 5 : [] (!!!)
  
  Testing grounds:
    [10,1,2,7,6,1,5], 8
    [1,1,2,5,6,7,10], 8
    
    candidates=[1:2,2:1,5:1,6:1,7:1,10:1], comb=[], combSum=0, target=8
    
    [1:1,2:1,5:1,6:1,7:1,10:1], comb=[1], combSum=1, target=8
    [2:1,5:1,6:1,7:1,10:1], comb=[1,1], combSum=2, target=8
    [5:1,6:1,7:1,10:1], comb=[1,1,2], combSum=4, target=8 -- [1,1,5] -- [1,1,6]
    [6:1,7:1,10:1], comb=[1,1,2,5], combSum=9, target=8  --- [5:1,7:1,10:1], comb=[1,1,2,6], combSum=9, target=8
                
    [5,1,2,1,6,7,10], 8
    
    [1:2,2:1,5:1,6:1,7:1,10:1], []
    
    [1:1,2:1,5:1,6:1,7:1,10:1], [1]
    [1:0,2:1,5:1,6:1,7:1,10:1], [1,1]
    
    [1,1,2,2,5,6,7,10], 8
         ^
*/
func combinationSum2(candidates []int, target int) [][]int {
  sort.Slice(candidates, func(i, j int) bool {
    return candidates[i] < candidates[j]
  })
  
  res := make([][]int, 0)
  backtrack(candidates, make([]int, 0, len(candidates)), target, &res)
  return res
}

func backtrack(candidates, comb []int, target int, res *[][]int) {
  if target == 0 {
    *res = append(*res, append([]int{}, comb...))
    return
  }
  
  if target < 0 {
    return
  }
  
  for i, v := range candidates {
    // Only go into the first occurencies of a number
    if i > 0 && candidates[i] == candidates[i-1] {
      continue
    }
    
    newTarget := target - v
    
    // Because everything is sorted, all the other numbers will ass well form newTarget < 0, so break preventivelly
    if newTarget < 0 {
      break
    }
    
    backtrack(candidates[i+1:], append(comb, v), newTarget, res)
  }
}






