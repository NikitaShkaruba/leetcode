/**
  Test cases:
    - A: 1
    - AA: 27 = 26 * 1 + 1 = 27
    - BA: 53 = 26 * 2 + 1 = 53
    - ZY: 701 = 26 * 26 + 25 = 676 + 25 = 701
    - ZZ: 26 * 26 + 26 = 702
    - AAA: 703 = 676 * 1 + 26 * 1 + 1 = 703
    - ZZZ: 676 * 26 + 26 * 26 + 26 = 17576 + 702 = 18278
    - AAAA: 676*26*1 + 676*1 + 26*1 + 1 = 17576 + X = 18279
    - FXSHRXW: ?
    
  Testing grounds:
  
    '?' = 'A' - 1 = 0
    'A' = '?' - 'A' = -'A' + 1 + 'A' = 1
    'B' = -'A' + 1 + 'B' = 2
    'Z' = -'A' + 1 + 'Z' = 

    'AA' = (- 'A' + 1 + 'A') * 27 = 26
    'BA' = (-'A' + 1 + 'B') * 27 = 54
    
    digit multiplication = 26^digitLocation, starting with 0

    A-----
     A----
      A--- 
       A-- 
        A- 26
         A - 1
         
     AAAA
        ^
        i := 4 - 1 = 3
        digitIndex := 4 - 1 - 3 = 0
         
 Solution:
    1. Brute force: O(n) time, O(1) space
*/
func titleToNumber(columnTitle string) int {
  num := 0
  
  base := 'A' - 1
  powerBase := 'Z' - base
  
  for i := len(columnTitle) - 1; i >= 0; i-- {
    val := int(rune(columnTitle[i]) - base)
    digitIndex := len(columnTitle) - 1 - i
    num += val * int(math.Pow(float64(powerBase), float64(digitIndex)))
  }
  
  return num
}