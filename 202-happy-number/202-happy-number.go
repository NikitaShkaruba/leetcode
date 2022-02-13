/**
  Test cases:
    - 19: true
    - 2: false
    - 1: true (!!!)

  Solutions:
    1. Cache: O(X) time, O(X) space, where X is the amount of iterations.
  
  Testing grounds:
    seen := {
      19,
      82,
      68,
      100
    }
    
    n := 19
    sum := 9*9 + 1*1 = 82
    divider := 100
    
    seen := {
      2,
      4,
      16,
      37,
      49 + 9 = 58,
      25 + 64 = 89,
      64 + 81 = 145,
      1 + 16 + 25 = 42,
    }
*/
func isHappy(n int) bool {
  seen := make(map[int]struct{})
  
  for n != 1 {
    if _, exists := seen[n]; exists {
      return false
    }
    seen[n] = struct{}{}
    
    n = numberDigitSquaresSum(n)
  }
  
  return true
}

/*
  12345
  
  12345 % 10 = 5
  math.Floor(12345 / 10) % 10 = 1234 % 10 = 4
  math.Floor(12345 / 100) % 10 = 123 % 10 = 3
  12345 / 10000 = 0.1
*/
func numberDigitSquaresSum(n int) int {
  sum := 0
  
  divider := 1
  for int(math.Floor(float64(n / divider))) != 0 {
    digit := int(math.Floor(float64(n / divider))) % 10
    sum += digit * digit
    divider *= 10
  }
  
  return sum
}