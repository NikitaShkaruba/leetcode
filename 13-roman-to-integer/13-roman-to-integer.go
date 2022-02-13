/* 
  Notes:
    - s validation?
    - Can be negative?
    - Nil string?
    - Next thing should exist
    
  Test cases:
    - MCMXCIV
    - MC
    - LVIII
    
  - M CM XC IV : 1000 + 900 + 90 + 4
         ^
    
  - MC : 1000 + 100
     ^
     
  - LVIII : 50 + 5 + 1 + 1 + 1
        ^
        
  - III
    ^
    
  Solutions:
    - O(n) time, O(1) space
*/
func romanToInt(s string) int {
  sum := 0
  
  for r := 0; r < len(s); r++ {
    curValue := romanLetterToInt(s[r])
    
    nextValue := 0
    if r != len(s) - 1 {
      nextValue = romanLetterToInt(s[r + 1])
    }
    
    if curValue >= nextValue {
      sum += curValue
    } else {
      sum += nextValue
      sum -= curValue
      r++
    }
  }
  
  return sum
}

func romanLetterToInt(r byte) int {
  symbolsMap := map[byte]int {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  }
  
  return symbolsMap[r]
}