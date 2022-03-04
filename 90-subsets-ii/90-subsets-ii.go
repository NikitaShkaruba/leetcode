/*

  Solutions:
    - Sort + add incrementally + hashset: O(n^2 * 2^n) time, O(2^n) space
    - Precompute amounts + add intcrementally: O(n^2 * 2^n) time, O(2^n) space
  
  Test caes:
    - [2,2,1] : [[],[1],[1,2],[1,2,2],[2],[2,2]]
    - [2,1,2] : [[],[1],[1,2],[1,2,2],[2],[2,2]]
    - [1,2,2] : [[],[1],[1,2],[1,2,2],[2],[2,2]]
    - [0] : [[], [0]] (!!!)
  
  Testing grounds:
  - [1,2,2] : [[],[1],[2],[1,2],[2,2],[1,2,2]]
       ^
     
    []
    [], [1] 
    [], [1], [2], [1, 2]
    [], [1], [2], [2,2], [1,2], [1, 2, 2]
    
    [1,2,2]: {1:1, 2:2}
    
    []
    [], [1] 
    
    [1,2,2]: {2:2, 1:1}
    []
    [], [2], [2, 2]
    [], [2], [2, 2], [1], [2, 1], [2, 2, 1]
  
    [1,2,2,2]
    {1:1, 2:3}
    
    []
    [], [1]
    [], [1], [2], [1, 2], [2, 2], [1, 2, 2], [2, 2, 2], [1, 2, 2, 2]
    
    nums: [2,1,2,2]
    counts: {1:1, 2:3}
    subsets: [], [1], [2], [1, 2], [2, 2], [1, 2, 2]
                                     ^ 
    num: 2, cnt: 3, subsetsIndex: 4, capturedSize: 4
    
    [4,4,4,1,4]
    counts: {4:4, 1:1}
    
    output:   [[],[4],[4,4],[4,4,4],[4,4,4,1],[1],[4,1],[4,4,1],[4,4,4,1],[4,4,4,1,1]]
    expected: [[],[1],[1,4],[1,4,4],[1,4,4,4],[1,4,4,4,4],[4],[4,4],[4,4,4],[4,4,4,4]]
    
*/

func subsetsWithDup(nums []int) [][]int {
  numCounts := map[int]int{}
  for _, v := range nums {
    if _, ok := numCounts[v]; !ok {
      numCounts[v] = 1
    } else {
      numCounts[v]++
    }
  }
  
  subsets := [][]int{
    []int{},
  }
  for num, cnt := range numCounts {
    jStart := 0
    for i := 0; i < cnt; i++ {
      curSubsetsLen := len(subsets)
      for j := jStart; j < curSubsetsLen; j++ {
        newSubset := make([]int, len(subsets[j]), len(subsets[j])+1)
				copy(newSubset, subsets[j])
				newSubset = append(newSubset, num)
        
				subsets = append(subsets, newSubset)
        
				jStart++
      }
    }
  }
  
  return subsets
}