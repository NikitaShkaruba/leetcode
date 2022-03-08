/*

  Solutions:
    1. HashMap + backtracking + remove duplicates: O(a lot) time, O(a lot) space
    2. HashMap + backtracking: O(n + n^(t/l)) time, O(n + t/l + n^(t/l)) space
  
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
*/
func combinationSum2(candidates []int, targetSum int) [][]int {
  candidateAmounts := make(map[int]int)
  for _, v := range candidates {
    if _, ok := candidateAmounts[v]; !ok {
      candidateAmounts[v] = 0
    }
    
    candidateAmounts[v]++
  }
  
  res := make([][]int, 0)
  backtrack(candidateAmounts, nil, 0, targetSum, &res)
  return res
}

func backtrack(candidateAmounts map[int]int, comb []int, combSum, targetSum int, res *[][]int) {
  if combSum == targetSum {
    *res = append(*res, append([]int{}, comb...))
    return
  }
  
  if combSum > targetSum {
    return
  }
  
  for v, c := range candidateAmounts {
    if c == 0 {
      continue
    }
    if len(comb) > 0 && v < comb[len(comb)-1] {
      continue
    }
    
    candidateAmounts[v]--
    backtrack(candidateAmounts, append(comb, v), combSum + v, targetSum, res)
    candidateAmounts[v]++
  }
}






