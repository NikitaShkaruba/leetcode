/*

  Solutions:
    - Brute force: O(n^3) time, O(1) space
    - Top-down + cachning: O(n^2) time, O(n^2) space
    - Middle pointer iteration: O(n^2), O(1) space (implemented) 
    
  Test cases:
    - "abc": 3
    - "aaa": 6
    - "a": 1
    - "abba": 6 (a, b, b, a, bb, abba)
  
  Testing grounds:
  
  abba
  a (0,0)
  ab (0,1)
  abb (0,2)
  abba (0,3)
  b (1,1)
  bb (1,2)
  bba (1,3)
  b (2,2)
  ba (2,3)
  a (3,3)
  
  abba
  0123
  
     abba
    /    \
  abb5     bba
  / \     / \
 ab2 bb3  bb  ba
 /\  /\  /\  /\
 a b b b b b b a
  
  if diff > 1 
    f(i,j) = f(i+1,j-1) + s[i] == [j] // if 
  else if diff == 1
    f(i,j) = s[i] == [j]
  else
    f(i,j) = true
  
  abba
    i
    l
     r
   
   count: 6
*/
func countSubstrings(s string) int {
  count := 0
  
  for i := 0; i < len(s); i++ {
    l := i
    r := i
    for l >= 0 && r < len(s) {
      if s[l] != s[r] {
        break
      }
      
      count++
      l--
      r++
    }
  }
  
  for i := 0; i < len(s) - 1; i++ {
    l := i
    r := i + 1
    for l >= 0 && r < len(s) {
      if s[l] != s[r] {
        break
      }
      
      count++
      l--
      r++
    }
  }
  
  return count
}