/*

  Solutions:
    - Brute force: O() time, O() space
    - Dynamic programming: O(n) time, O(n) space
    - Dynamic programming + store last 2 maxes: O(n) time, O(1) space
  
  Test cases:
    - [2,3,2]: 3
    - [1,2,3,1]: 4
    - [1,2,3]: 3
    - [5]: 5 (!!!)
  
  Testing grounds:
    - [2,3,2]
      [2,3]: 3
    - [2,2]: 2
    
    
    - [1,2,3,1]: 4
      
      [1,2,3]
      [1,2,4]: 4
      
      [1,1,2]
      [1,1,3]: 3
      
    - [1,2,3]
    
      [1,2]
      [1,2]
      
      [3,1]
      [3,3]: 3
      
    - [1,2,3,1]
             ^
    prev := 4
    prevPrev := 2
    cur = 4
    
    prev := 1
    prevPrev := 1
    cur
*/
func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return maxInt(nums[0], nums[1])
    }

    max1 := simpleRob(nums[:len(nums)-1])
    max2 := simpleRob(nums[1:])

    return maxInt(max1, max2)
}

func simpleRob(nums []int) int {
    prevPrevValue := nums[0]
    prevValue := maxInt(nums[0], nums[1])

    for i := 2; i < len(nums); i++ {
        curValue := maxInt(nums[i]+prevPrevValue, prevValue)

        prevPrevValue = prevValue
        prevValue = curValue
    }

    return prevValue
}

func maxInt(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}