/*

  Solutions:
    - Brute force: O(n*len(n)) time, O(1) space
    - Dynamic programming: O(n) time, O(n) space
  
  Test cases:
    - 0: [0]
    - 5: [0, 1, 1, 2, 1, 2]
  
  Testing grounds:

    - 0 --> 0
    - 1 --> 1
    - 2 --> 10
    - 3 --> 11
    - 4 --> 100
    - 5 --> 101
    
    1100100101 = 100000000 + 100000000 + 100000 + 100 + 1
    
    1
    10
    11
    100
    101
    110 = 100 + 10
    111 = 100 + 11
    1000 % 1000

    [0, 1*, 1*, 2, 1*, 2, 2, 3, 1*, 2, 2, 3, 2(12), 3(13), 3(14), 4(15), 1*]
                   p            n
*/

func countBits(n int) []int {
  if n == 0 {
    return []int{0}
  }
  
  cache := make([]int, n+1)
  cache[0] = 0
  cache[1] = 1
  
  prevPowerOfTwo := 1
  nextPowerOfTwo := 2
  
  for i := 2; i <= n; i++ {
    if i == nextPowerOfTwo {
      cache[i] = 1
      prevPowerOfTwo = nextPowerOfTwo
      nextPowerOfTwo *= 2
      continue
    }
    
    diff := i - prevPowerOfTwo
    cache[i] = 1 + cache[diff]
  }
  
  return cache
}