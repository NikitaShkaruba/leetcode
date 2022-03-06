/*

  Solutions:
    1. Sort + backtrack: O(n^3) time, O(n!) space
    2. Count all + bracktrack: O(n^2 * n!) time, O(n) space
    
  Test cases:
    - [1] : [[1]]
    - [1,1,2] : [[1,1,2],[1,2,1],[2,1,1]]
    - [1,1,1,1,2,2]
    - [1,2,3] : [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
  
  Testing grounds:

          [1,2,3], []
          /          \
    [2, 3], [1] -- [1, 3], [2]  -- [1, 2], [3]
      |
  [], [1, 2, 3] -- ...
  
          [1,2,3,2], []
          /                \
  [2,3,2], [1] -- [1,3,2], [2] -- [1,2,2], [3] -- [1,2,3], [2]
  
                       |                              |
                    [], [2,1,3,2]                 [], [2,1,3,2]
                    
              [1,2,2,3], []
          /                \
  [2,2,3], [1] -- [1,2,3], [2] -- [1,2,3], [2] -- [1,2,2], [3]
                      |               |
                     ...             ...
                      |               |
                [], [2,1,3,2]     [], [2,1,3,2]
                
      [1], []
         |
      [], [1]
      
                  [1,2,1,3]
       [1,2], [1]  -- [x] -- [1,1], [2]
  [2], [1,1]  -- [1], [1,2]    [1], [2, 1] -- [x]
  
    [1,2,1,3,1]
    [1,1,1,2,3], []
    /
    [1,1,2,3], [1] -- [1,1,1,3], [2] -- [1,1,1,2], [3]
    
    [1,1,1,2,3], []
    
    []
*/

func permuteUnique(nums []int) [][]int {
  numAmounts := make(map[int]int)
  for _, v := range nums {
    if _, ok := numAmounts[v]; !ok {
      numAmounts[v] = 0
    }

    numAmounts[v]++
  }

  res := make([][]int, 0)
  backtrack(numAmounts, []int{}, len(nums), &res)

  return res
}

func backtrack(numAmounts map[int]int, perm []int, finishedPermLen int, res *[][]int) {
  if len(perm) == finishedPermLen {
    permCopy := make([]int, len(perm))
    copy(permCopy, perm)
    *res = append(*res, permCopy)
    return
  }

  for num := range numAmounts {
    if numAmounts[num] == 0 {
      continue
    }

    numAmounts[num]--
    backtrack(numAmounts, append(perm, num), finishedPermLen, res)
    numAmounts[num]++
  }
}