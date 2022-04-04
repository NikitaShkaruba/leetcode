/*

  Solution:
    - Brute force: O(k^n) time, O(n) space (implemented)
  
  Test cases:
    - [4,3,2,3,5,2,1], k = 4: true
    - [1,2,3,4], k = 3: false
  
  Testing grounds:
    [4,3,2,3,5,2,1], k = 4
    subsetSum = sum(nums) / k = 5
    
    [4,3,2,3,5,2,1]
     i           j
     
    f(nums, subsetSum, currentSum, sumsNeeded)
    
    nums = [5]
            i
    subsetSum = 5
    currentSum = 0
    sumsNeeded = 0
    
    exit when sumsFound == k-1
    
*/
func canPartitionKSubsets(nums []int, k int) bool {
  if k == 1 {
    return true
  }
  
  numsSum := sliceSum(nums)
  if numsSum % k != 0 {
    return false
  }
  
  seenNums := make([]bool, len(nums))
  targetSum := numsSum / k
  return backtrack(0, nums, seenNums, targetSum, 0, k-1)
}

func backtrack(i int, nums []int, seenNums []bool, targetSum, currentSum, sumsNeeded int) bool {
  if sumsNeeded == 0 {
    return true
  }
  
  if currentSum > targetSum {
    return false
  }
  if currentSum == targetSum {
    return backtrack(0, nums, seenNums, targetSum, 0, sumsNeeded-1)
  }
  
  for j := i; j < len(nums); j++ {
    if seenNums[j] {
      continue
    }
    seenNums[j] = true
    
    if backtrack(j + 1, nums, seenNums, targetSum, currentSum + nums[j], sumsNeeded) {
      return true
    }
    
    seenNums[j] = false
  }
  
  return false
}

func sliceSum(s []int) int {
  sum := 0
  
  for _, v := range s {
    sum += v
  }
  
  return sum
}