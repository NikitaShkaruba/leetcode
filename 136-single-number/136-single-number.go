/**
  Test cases:
    - [2, 2, 1]
    - [4, 1, 2, 1, 2]
    - [1] (!!!)
    
  Solutions:
    - Brute force: O(n^2) time, O(1) space
    - Cache approach: O(n) time, O(n) space
    - XOR: O(n) time, O(1) space
    
  Testing grounds:
  - [4, 1, 2, 1, 2]
    4 + 1 + 2 + 1 + 2 = 10 / 5 = 2
  
    cache = {
      4: true
    }
*/
func singleNumber(nums []int) int {
  xor := 0
  
  for _, v := range nums {
    xor ^= v
  }
  
  return xor
}