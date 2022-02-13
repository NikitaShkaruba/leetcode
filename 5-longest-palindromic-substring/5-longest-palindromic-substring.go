/*

  Test cases:
    - babad: bab
    - babbad: abba
    - abcdef: a
    - cbbd: bb
    - aaaaa: a
    - abcdd (!!!)
    - a: a (!!!)
  
  Solutions:
    - Brute force: O(n^2) time, O(1) space
  
  Testing grounds:
    - babad
         l^r 
      
      best: bab
      l^^r
      3 - 1 + 1 = 3
      
    - bcabbacd
       l ^^ r
      
      best: cabbac
      
    - babbad
        i
       l
          r
  
*/
func longestPalindrome(s string) string {
  result := ""
  
  for i := 0; i < len(s); i++ {
    // check if i as pivot
    l := i
    r := i
    for l >= 0 && r < len(s) && s[l] == s[r] {
      if r - l + 1 > len(result) {
        result = s[l:r+1]
      }  
      
      l--
      r++
    }
    
    // check if i and i+1 is pivot
    l = i
    r = i + 1
    for l >= 0 && r < len(s) && s[l] == s[r] {
      if r - l + 1 > len(result) {
        result = s[l:r+1]
      }  
      
      l--
      r++
    }
  }
  
  return result
}