func mySqrt(x int) int {
  for i := 1; i <= x; i++ {
    if i*i == x {
      return i 
    }

    if i*i > x {
      return i-1
    }
  }

  return 0
}