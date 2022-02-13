/*
      _
     __
    ___
   ____
  _____
 ______
_______

n = 5
[1, 1, 2, 3, 5]
   -  -  -
   ----  -
   -  ----
*/

var cache map[int]int = make(map[int]int)

func climbStairs(n int) int {
  if n <= 1 {
    return 1
  }
  
  if _, exists := cache[n]; exists {
    return cache[n]
  }
  
  cache[n] = climbStairs(n-1) + climbStairs(n-2)
  
  return cache[n]
}