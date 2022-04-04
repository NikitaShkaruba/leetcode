/*
  
  Solutions:
    - Brute force: O(2^n) time, O(n) space
    - Brute force + caching (Top-down): O(n*) time, O(2^n + n) space (will be faster with big nums length)
    - Tabulation (Bottom-up): O() time, O() space
  
  Test cases:
    - [1,5,11,5]: 11: true
    - [1,2,3,5]: 6, 5: false
    - [1]: false

  Testing ground:

    
                                 f(leftSum, rightSum, nums)
                                    0, 0, [1,5,11,5]
                     1, 0, [5,11,5]                     0, 1, [5,11,5]
             6, 0, [11,5]     1, 5, [11,5]      5, 1, [11,5]        0, 6, [11,5]
    17, 0, [5]     6, 11, [5]  
                   11, 11, []
                     

      0 1 5 11 5
    0 x 1 6 17 22
    1 1 x 5       
    5 6   x     
   11 17    x   
    5 22       x

    [1,5,11,5] target = 11
    [1]
*/
func canPartition(nums []int) bool {
  numsSum := sliceSum(nums)
  if numsSum % 2 != 0 {
    return false
  }
  
  cache := map[[2]int]struct{}{}
  return helper(nums, numsSum / 2, cache)
}

func helper(nums []int, target int, cache map[[2]int]struct{}) bool {
  if target == 0 {
    return true
  }
  if len(nums) == 0 {
    return false
  }
  
  cacheKey := [2]int{len(nums), target}
  if _, ok := cache[cacheKey]; ok {
    return false
  }
  
  newNums := nums[1:]
  res := helper(newNums, target, cache) || helper(newNums, target - nums[0], cache)
  
  cache[cacheKey] = struct{}{}
  
  return res 
}

func sliceSum(s []int) int {
  sum := 0
  
  for _, v := range s {
    sum += v
  }
  
  return sum
}