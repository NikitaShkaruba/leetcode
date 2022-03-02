/*
  
  Solutions:
    - Count symbols: O(n) time, O(n) space
  
  Test cases:
    - c : true (c:1)
    - cc : true (c:2)
    - cca : true (c:2, a:1)
    - acac : true (a:2, c:2)
    - cabc : false (a:1, b:1, c:1)
    - cccc : true (c:4)
    - ccccc : true (c:5)
    - cccccccccccccccccccccccccc: true
    - carerac : true
  
  Testing grounds:
    carerac
    {
      c:2,
      a:2,
      r:2
      e:1
    }
*/
func canPermutePalindrome(s string) bool {
  symbolAmounts := make(map[byte]int)
  
  for _, v := range []byte(s) {
    if _, exists := symbolAmounts[v]; !exists {
      symbolAmounts[v] = 0
    }
    
    symbolAmounts[v]++
  }
  
  oddsAmount := 0
  for _, v := range []byte(s) {
    amount, notProcessed := symbolAmounts[v]
    if !notProcessed {
      continue
    }
    
    if amount % 2 == 1 {
      oddsAmount++
      if oddsAmount > 1 {
        return false
      }
    }
    
    delete(symbolAmounts, v)
  }
  
  return true
}