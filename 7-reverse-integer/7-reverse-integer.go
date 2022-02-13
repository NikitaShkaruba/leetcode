/*

  Test cases:
    - 123: 321
    - -123: -321
    - 120: 21
    - 102: 201
    - 111111119: 0 (!!!)
    - 0: 0 (!!!)
  
  Solutions:
    1. Brute force: O(n) time, O(1) space
  
  Testing grounds:
    -120: -21
    
    1 + 20 + 0
    
    120 / 1 % 10 = 0
    120 / 10 % 10 = 2
    120 / 100 % 10 = 1
    
    120
      ^
    
    divider: 0
    multiplier: 1000
    reversedX: 1 + 20 = 21 * -1 = -21

    120 / 1 = 120
    120 / 10 = 12
    120 / 100 = 1 (+)
    120 / 1000 = 0
    divider: 1
    
    99 / 1 = 99
    99 / 10 = 9.9
    99 / 100 = 0.99
  
    9 / 1 = 9
    9 / 10 = 0.9
*/
func reverse(x int) int {
  isXNegative := x < 0
  x = int(math.Abs(float64(x)))
  
  divider := maxDivider(x)
  multiplier := 1
  reversedX := 0
  
  for divider >= 1 {
    digit := int(math.Floor(float64(x / divider))) % 10
    reversedX += digit * multiplier
    
    if reversedX > math.MaxInt32 {
      return 0
    }
    
    divider /= 10
    multiplier *= 10
  }
  
  if isXNegative {
    reversedX *= -1
  }
  
  return reversedX
}

func maxDivider(x int) int {
  divider := 1
  
  for int(math.Floor(float64(x / (divider * 10)))) > 0 {
    divider *= 10 
  }
  
  return divider
}