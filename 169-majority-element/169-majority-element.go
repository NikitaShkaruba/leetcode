func majorityElement(nums []int) int {
  amount := 0
  result := 0
  
  for i := 0; i < len(nums); i++ {
      if amount == 0 {
          result = nums[i]
          amount++
      } else if nums[i] == result {
          amount++
      } else {
          amount--
      }
  }
  
  return result
}