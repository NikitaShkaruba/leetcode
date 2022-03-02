/*

  Solutions:
    - DFS: O(n^n * n*logn) time, O(2^n) space
    - Iterate and mutate with current subsets: O(n*2^n) time, O(2^n) space
  
  Test cases:
    - [1,2,3] : [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
    - [0] : [[],[0]]
  
  Testing grounds:
    [1,2,3]
    
    [] - [1] - [1, 2] - [1, 2, 3]
             - [1, 3]
       - [2] - [2, 3]
       - [3]
     
     nums = [1,2,3]
     subsets = [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]
     Tsubset = [[], [1], [2], [1, 2]]
     
     nums = [1,2]
     subsets = [[], [1], [2], [1, 2], [3]]
     
     nums = [0]
             ^
     subsets = [[], [0]]
*/

func subsets(nums []int) [][]int {
  subsets := [][]int{
    []int{},
  }
  
  for _, num := range nums {
    for _, subset := range subsets {
      newSubset := make([]int, len(subset) + 1)
      copy(newSubset, subset)
      newSubset[len(newSubset)-1] = num
      
      subsets = append(subsets, newSubset)
    }
  }
  
  return subsets
}