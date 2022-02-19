/*

  Solutions:
    - Just nums (Brute force): 
      - Creation: O(n) time, O(n) space 
      - SumRange: O(n) time, O(1) space
    - Sums: 
      - Creation: O(n) time, O(n) space 
      - SumRange: O(1) time, O(1) space
  
  Test cases:
    Creation:
      - [-2, 0, 3, -5, 2, -1]
      - [1] (!!!)
    SumRange:
      - [0, 2] (!!!)
      - [1, 2]
  
  Testing grounds:

    nums:      [-2,  0,  3, -5,  2, -1]        
    sums:      [-2, -2,  1, -4, -2, -3]
                     l       r    
    nums:      [-2] 
    sums:      [-2]
                 ^
*/

type NumArray struct {
  sums []int
}


func Constructor(nums []int) NumArray {
  res := NumArray{
    sums: make([]int, 0, len(nums)),
  }
  
  curSum := 0
  for _, v := range nums {
    curSum += v
    res.sums = append(res.sums, curSum)
  }
  
  return res
}


func (this *NumArray) SumRange(left int, right int) int {
  if left == 0 {
    return this.sums[right]
  }
  
  return this.sums[right] - this.sums[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */