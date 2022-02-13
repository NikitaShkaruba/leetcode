func isSubsequence(t string, s string) bool {
  if len(t) == 0 {
    return true
  }
  
  j := 0
  
  for i := range s {
    if s[i] == t[j] {
      j++
      
      if j == len(t) {
        return true
      }
    }
  }
  
  return false
}