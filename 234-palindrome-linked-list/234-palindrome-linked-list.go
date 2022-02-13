/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
  Test cases:
    - 1 => 2 => 2 => 1 : true
    - 1 => 2 => 3 => 2 => 1 : true
    - 1 => 2 => 3 => 3 => 1 : false
    - 1 => 2: false
    - 1 : true (!!!)
 
  Solutions:
    1. Create a string: O(n) time, O(n) space 
    2. Reverse the last half: O(n) time, O(1) space 
  
  Testing grounds:
    
    1 => 2 => 2 => 1 : "1221"
    
    1 -> 2 -> 2 -> 1
    1 -> 2 -> 2 <- 1
         ^^
         
    - 1 => 2 <= 2 => 1 : true (4)
           ^    ^
    - 1 => 2 <= 3 <= 2 <= 1 : true (5) 2
           ^         ^
           
    1 <- 2 <- 2 <- 1 -> nil
    
    1. Compute len
    2. Iterate and reverse after len / 2, remember the last element
    3. Iterate for len/2, check for equality
    
    
    1 => 2 => 3 <=> 2 <=  1
              lr
              
    1 => 2 <=> 2 <= 1
            lr
            
    1 => nil
    lr
    
    halfListLen := 0
    j := 1
 */
func isPalindrome(head *ListNode) bool {
  halfListLen := int(math.Floor(float64(listLen(head) / 2)))
  
  prev := head
  cur := head.Next
  i := 0
  for cur != nil {
    if i >= halfListLen {
      tmp := cur.Next
      cur.Next = prev
      prev = cur
      cur = tmp
    } else {
      prev = cur
      cur = cur.Next
    }
    
    i++
  }
  
  left := head
  right := prev
  
  for j := 0; j < halfListLen; j++ {
    if left.Val != right.Val {
      return false
    }
    
    left = left.Next
    right = right.Next
  }
  
  return true
}

func listLen(head *ListNode) int {
  listLen := 0
  
  cur := head
  for cur != nil {
    listLen++
    cur = cur.Next
  }
  
  return listLen
}















