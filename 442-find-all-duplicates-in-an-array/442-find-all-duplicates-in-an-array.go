/*

  Solutions:
    - Brute force: O(n^2) time, O(1) space
    - Hashset cache: O(n) time, O(n) space 
    - Negative marks: O(n) time, O(1) space
  
  Test cases:
    - [4, 3, 2, 2, 3]: [3, 2]
    - [1, 2, 3]: [] (!!!)
    - [1]: [] (!!!)
  
  Testing grounds:
    [-1,-1,2]
          ^
*/
func findDuplicates(nums []int) []int {
  duplicates := []int{}
  
  for i := range nums {
    v := int(math.Abs(float64(nums[i])))
    
    if nums[v-1] < 0 {
      duplicates = append(duplicates, v)
    } else {
      nums[v-1] = -nums[v-1]
    }
  }
  
  return duplicates
}