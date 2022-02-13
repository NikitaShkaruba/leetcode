func isValid(s string) bool {
  stack := make([]rune, 0)
  m := map[rune]rune {
    ')': '(',
    ']': '[',
    '}': '{',
  }
  
  for _, c := range s {
    switch c {
    case '(', '{', '[':
      stack = append(stack, c)
    case ')', '}', ']':
      if len(stack) == 0 || stack[len(stack) - 1] != m[c]{
        return false
      }
      stack = stack[:len(stack)-1]
    }
  }

  return len(stack) == 0
}