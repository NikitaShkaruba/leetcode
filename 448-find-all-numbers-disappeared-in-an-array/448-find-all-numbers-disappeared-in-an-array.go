/*

  Solutions:
    1. Brute force: O(n) time, O(n) space
    2. Mark with minuses: O(n) time, O(1) space
  
  Test cases:
    - [4, 3, 2, 7, 8, 2, 3, 1]
    - [3, 3, 2] => [1]
    - [3, 3, 3, 3] => [1, 2, 4]
    - [1] (!!!)
  
  Testing grounds:
    nums: [3, 3, 2]
    set: {
      1: true,
    }
    
    nums: [-4, -3, -2, -7, 8, 2, -3, -1]
                              ^
    res: [5, 6]                     
               
    iterate and mark. Mark index is index - 1
    iterate and store indexes to the result array
*/

func findDisappearedNumbers(nums []int) []int {
  for _, v := range nums {
    indexToMark := int(math.Abs(float64(v))) - 1
    
    if nums[indexToMark] > 0 {
      nums[indexToMark] *= -1
    }
  }
                       
  res := make([]int, 0)
  for i, v := range nums {
    if v > 0 {
      res = append(res, i + 1)
    }
  }
  
  return res
}