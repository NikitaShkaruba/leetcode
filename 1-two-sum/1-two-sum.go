func twoSum(nums []int, target int) []int {
  complimentsSet := make(map[int]int)
  
  for i := range nums {
    compliment := target - nums[i]

    if j, exists := complimentsSet[compliment]; exists {
      return []int{i, j}
    }
    
    complimentsSet[nums[i]] = i
  }
  
  // Wouldn't be hit once
  return []int{}
}