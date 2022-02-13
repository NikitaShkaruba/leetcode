func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
   
    for i < j {
        if !isValid(s[i]) {
            i++
            continue
        }
        if !isValid(s[j]) {
            j--
            continue
        }
      
        if !strings.EqualFold(string(s[i]), string(s[j])) {
            return false
        }
      
        i++
        j--
    }
  
    return true
}

func isValid(a byte) bool {
  return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}