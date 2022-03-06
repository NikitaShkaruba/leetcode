/*

  Solutions:
    1. Sort + backtrack: O(n^3) time, O(n!) space
    2. Backtrack + hashtable: O(n^3) time, O(n!) space
    
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
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)
	backtrack(nums, []int{}, &res)

	return res
}

func backtrack(nums []int, perm []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, perm...))
		return
	}

	for i := range nums {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}

		numsWithoutI := make([]int, 0, len(nums)-1)
		for j, v := range nums {
			if j == i {
				continue
			}
			numsWithoutI = append(numsWithoutI, v)
		}

		backtrack(numsWithoutI, append(perm, nums[i]), res)
	}
}