/*

  Solutions:
    - Brute force: O(n^n) time, O(n) space
    - Sorting and iterating: O(n*logn) time, O(n) space
    - Merge intervals: O(n) time, O(n) space, hard to implement
    - Count and put all to set: O(n) time, O(n) space
  
  Test cases:
    - [1,2,3,4,2]: 1
    - [1,1,1,1,1]: 1
    - [1,3,5,6,8]: 1
    - [100,0,200,1,-1,2] : 4
    - [1] : 1 
    - [] : 0 (!!!)
  
  Testing grounds:
    - [100,0,200,1,-1,2]
                      ^                    
      {
        100: [100, 100],
        200: [200, 00],
        -1: [-1, 2],
        2: [-1, 2]
      }
      
    - [1,2,3,4,2,3,2,1,4]: 1
       ^
    
    {
      1: &[1,4]
      4: &[1,4]
      2: &[2,3]
      3: &[2,3]
    }

*/

func longestConsecutive(nums []int) int {
  longest := 0
  
  set := make(map[int]struct{})
  for _, v := range nums {
    set[v] = struct{}{}
  }
  
  for _, v := range nums {
    if _, exists := set[v-1]; exists {
      continue
    }
    
    curLength := 0
    num := v
    for going := true; going; _, going = set[num] {
      curLength++
      num++
    }
    
    if curLength > longest {
      longest = curLength
    }
  }
 
  return longest
}
