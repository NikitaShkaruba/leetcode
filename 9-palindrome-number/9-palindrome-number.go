/*
  1034301 => "1034301"
  1034301
  l     r
  
  l1: floor(number / 1000000) % 10
  r1: floor(number / 1) % 10 
  
  l1: floor(number / 100000) % 10
  r1: floor(number / 10) % 10 
  
  Solutions:
    - Convert to string: O(n) time, O(len(n)) space
    - Two pointers: O(n) time, O(1) space 
*/
func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }
  
  leftDiv := computeLeftDiv(x)
  rightDiv := 1
  
  for leftDiv > rightDiv {
    leftNum := int(math.Floor(float64(x / leftDiv))) % 10
    rightNum := int(math.Floor(float64(x / rightDiv))) % 10
    
    if leftNum != rightNum {
      return false
    }
    
    leftDiv /= 10
    rightDiv *= 10
  }
  
  return true
}

// 4 => 1
// 11 => 
// 10 => 
// 99 =>
// 100
func computeLeftDiv(x int) int {
  div := 1
  
  for div * 10 <= x {
    div *= 10
  }
  
  return div
}
