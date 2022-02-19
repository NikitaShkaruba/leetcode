/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
  
  Solutions:
  
  Test cases:
    - 1 -> 2 -> 3, 1
    - 1 -> 2 -> 3, 2
    - 1 -> 2 -> 3, 3
    - 1 -> 1 -> 1, 1
    - 1, 1 (!!!)
    -  (!!!)
    
  Testing grounds:
    - 1 -> 3, 2
      p         c
      h
      
    - 1 -> 2 -> 3, 2
           p    c
  
    - 1 -> 2 -> 3, 1 (+ move the head optionally)
  p        c
  
    - 1 -> 1 -> 1, 1
                    c
                    h  
          
    - 1, 1 (!!!)
      c
      
    -  (!!!)
    
    - 1 -> 2 -> 3, 1
           c
           h
*/

func removeElements(head *ListNode, val int) *ListNode {
  var prev *ListNode = nil
  cur := head
  
  for cur != nil {
    if cur.Val != val {
      prev = cur
      cur = cur.Next
      continue
    }
    
    if prev == nil {
      head = head.Next
    } else {
      prev.Next = cur.Next
    }
    
    cur = cur.Next
  }
  
  return head
}













