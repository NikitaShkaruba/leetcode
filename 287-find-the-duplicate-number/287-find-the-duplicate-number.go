/*

  Solutions:
    1. Seen set: O(n) time, O(n) space
    2. Seen set in nums: O(n) time, O(1) space, but modifies the input array
    3. Double iteration: O(n^2) time, O(1) space
    4. Double iteration with Binary search: O(n^2) time, O(1) space
  
  Test cases:
    - [1,3,4,2,2]: 2
    - [3,1,3,4,2]: 3
    - [2,2,2,2,2]
    - [1, 1]: 1 (!!!)
  
  Testing grounds:
    - [1,3,4,2,2], 
    Got sum: 12
    Target sum: 10
    Result: 2
    
    - [3,1,3,4,2]
    Got sum: 13
    Target sum: 10
    Result: 3

    - [2,2,2,2,2]
    Got sum: 10
    Target sum: 10
    
    Target Sum: 10 
    Min Sum: 5
    Max Sum: 20
    
    - [1,2,3,4,4]: 14
    - [1,2,3,4,4]: 14
    - [1,4,4,4,4]: 17
    - [4,4,4,4,4]: 20
    
    - [1,2,3,4,2]: 12
    - [1,2,3,2,2]: 10
    - [1,2,2,2,2]: 9
    - [2,2,2,2,2]: 10
*/
func findDuplicate(nums []int) int {
  tortoise := nums[0]
  hare := nums[0]
  
  for ok := true; ok; ok = tortoise != hare {
    tortoise = nums[tortoise]
    hare = nums[nums[hare]]
  }
  
  hare = nums[0]
  for tortoise != hare {
    tortoise = nums[tortoise]
    hare = nums[hare]
  }
  
  return hare
}