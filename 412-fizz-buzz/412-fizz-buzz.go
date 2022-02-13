func fizzBuzz(n int) []string {
  result := []string{}
  
  for i := 1; i <= n; i++ {
    s := ""
    
    if i % 3 == 0 {
      s += "Fizz"
    }
    if i % 5 == 0 {
      s += "Buzz"
    }
    if s == "" {
      s = strconv.Itoa(i)
    }
    
    result = append(result, s)
  }
  
  return result
}