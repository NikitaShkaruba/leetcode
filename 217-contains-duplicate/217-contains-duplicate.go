/*

  Test cases:
    - [1, 2, 3, 1]: true
    - [1, 2, 3, 4]: false
    - [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]: true
    - [1]: false
  
  Solutions:
    1. nums cache: O(n) time, O(n) space

  Testing grounds:
  
*/
func containsDuplicate(nums []int) bool {
  seen := make(map[int]struct{})
  
  for _, n := range nums {
    if _, exists := seen[n]; exists {
      return true
    }
    seen[n] = struct{}{}
  }
  
  return false
}