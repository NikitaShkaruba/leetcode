/*

  Solutions:
    1. Stack solution: O(n + t) time, O(n + t) space
    2. Two pointers: O(min(n, t)) time, O(1) space
  
  Test cases:
    - "ab#c", "ad#c": true
    - "a#c", "b": false
    - '#######', '123': false (!!!)
  
  Testing grounds:

    #ab#c
      ^
    ad#c###
     ^
    
    ###
   ^
   
   ab##
  ^
   c#d#
      ^
      
   ab#c, 0
   ^
   ad#c, 0
   ^
*/
func backspaceCompare(s string, t string) bool {
  i := len(s) - 1
  j := len(t) - 1
  
  for i >= 0 || j >= 0 {
    toSkipI := 0
    for i != -1 && (s[i] == '#' || toSkipI > 0) {
      if s[i] == '#' {
        toSkipI++
      } else {
        toSkipI--
      }
      
      i--
    }
    
    toSkipJ := 0
    for j != -1 && (t[j] == '#' || toSkipJ > 0)  {
      if t[j] == '#' {
        toSkipJ++
      } else {
        toSkipJ--
      }
      
      j--
    }
    
    if i == -1 && j == -1 {
      return true
    }
    if i == -1 || j == -1 {
      return false
    }
    
    if s[i] != t[j] {
      return false
    }
    
    i--
    j--
  }
  
  return true
}