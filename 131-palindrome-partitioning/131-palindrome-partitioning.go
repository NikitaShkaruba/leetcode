/*

  Solutions:
    - Backtrack: O(n * 2^n) time, O(n * 2^n) space
  
  Test cases:
    - "aabaa": [["a","aba",a],["a","a","b","a","a"], ["a", "a", "b", "aa"], ["aa","b", "aa"],["aa","b", "a","a"], "aabaa"]
    - "aab": [["a","a","b"],["aa","b"]]
    - "a": [["a"]]
  
  Testng grounds: 
    [],abba
      [a], bba
        - [a, b], ba
          - [a, b, b], a
            - [a, b, b, a] (V)
          - [a, b, ba] (X)
        - [a, bb], a
          - [a, bb, a] (V)
        - [a, bba] (X) 
      [ab], ba (X)
      [abb], a (X)
      [abba], (V)
  
*/
func partition(s string) [][]string {
  res := [][]string{}
  backtrack(s, []string{}, make(map[string]bool), &res)
  return res
}

func backtrack(s string, comb []string, cache map[string]bool, res *[][]string) {
  if len(s) == 0 {
    *res = append(*res, append([]string{}, comb...)) 
    return
  }
  
  for i := 1; i <= len(s); i++ {
    subS := s[0:i]
    
    isPali, ok := cache[subS]
    if !ok {
      isPali = isPalindrome(subS)
      cache[subS] = isPali
    }
        
    if !isPali {
      continue 
    }
    
    backtrack(s[i:], append(comb, subS), cache, res)
  }
}

func isPalindrome(s string) bool {
  l := 0
  r := len(s) - 1
  
  for l < r {
    if s[l] != s[r] {
      return false
    }
    
    l++
    r--
  }
  
  return true
}