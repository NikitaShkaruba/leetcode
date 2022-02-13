/*

  Test cases: 
    - abcabcbb => 3
    - bbbbbb => 1
    - pwwkew => 3
    - abcdef => 6
    - "" => 0 (!!!)
    - "a" => 1 (!!!)
  
  Solutions:
    - Brute force: O(n^2) time, O(n) space
    - Sliding window: O(n) time, O(n) space
  
  Testing ground:
  
    pwwkew
       l
          r
    
    set := {
      k,
      e,
      w
    }
  
    bestLen := 3
    
    
    bbbbbb
      l
       r

*/
func lengthOfLongestSubstring(s string) int {
  if len(s) == 0 {
    return 0
  }
  
  usedChars := map[byte]struct{}{
    s[0]: struct{}{},
  }
  bestLen := 1
  
  l := 0
  r := 1
  for r < len(s) {
    _, containsUsedChar := usedChars[s[r]]
    for containsUsedChar {
      delete(usedChars, s[l])
      l++
      
      _, containsUsedChar = usedChars[s[r]]
    }
    
    usedChars[s[r]] = struct{}{}
    r++
    
    if len(usedChars) > bestLen {
      bestLen = len(usedChars)
    }
  }
  
  return bestLen
}