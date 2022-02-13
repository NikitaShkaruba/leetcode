func sortedSquares(nums []int) []int {
  sortedSquares := make([]int, len(nums))
  
  l := 0
  r := len(nums) - 1
  c := len(nums) - 1
  
  for l <= r {
    var biggerSquare int
    if nums[l]*nums[l] > nums[r]*nums[r] {
      biggerSquare = nums[l]*nums[l]
      l++
    } else {
      biggerSquare = nums[r]*nums[r]
      r--
    }
    
    sortedSquares[c] = biggerSquare
    c--
  }
  
  return sortedSquares
}