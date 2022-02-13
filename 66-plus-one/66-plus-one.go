/**
  Test cases:
    - [1, 2, 3]
    - [4, 3, 2, 1]
    - [9, 9, 9]
    - [9]
    
  Testing grounds:
    - [1, 2, 4]
    - [9, 9, 9] => [1, 0, 0, 0]
       ^
    
  Solutions:
    - Bruteforce: O(n) time, O(n) space
*/
func plusOne(digits []int) []int {
  for i := len(digits) - 1; i >= 0; i-- {
    if digits[i] != 9 {
      digits[i]++
      return digits
    }
    
    digits[i] = 0
    if i == 0 {
      digits[0] = 1
      digits = append(digits, 0)
    }
  }
  
  return digits
}