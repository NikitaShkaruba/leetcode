func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
      return false
  }
  
  chars := make(map[rune]int)
  
  for _, v := range s {
    chars[v]++
  }
  
  for _, v := range t {
    chars[v]--  
    if chars[v] == 0 {
      delete(chars, v)
    }
  }
  
  return len(chars) == 0
}