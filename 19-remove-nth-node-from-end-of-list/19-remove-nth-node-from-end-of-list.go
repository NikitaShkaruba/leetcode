/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*

  Solutions:
    - PreDel iterator: O(l) time, O(1) space, where l is the length of the list
  
  Test cases:
    - 1, n = 1: nil (!!!)
    - 1 -> 2 -> 3, n = 3: 2 -> 3
    - 1 -> 2 -> 3, n = 1: 1 -> 2
    - 1 -> 2 -> 3, n = 2: 1 -> 3
  
  Testing grounds:
  
    1 -> 2 -> 3 -> 4 -> 5, n = 2 
             p          c
    
    1
    c
    
    1 -> 2 -> 3 -> 4 -> 5, n = 2
    p         c
    
    currentNodeIndex = 2
    
    remember preDel when it is equal to n
    iterate
  
    after the cycle
    then either remove p+1, or if p = nil, remove head. Don't forget to return head.next or nil

*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  cur := head
  var preDel *ListNode
  currentNodeIndex := 0
  
  for cur.Next != nil {
    cur = cur.Next
    currentNodeIndex++
    
    if preDel == nil && currentNodeIndex == n {
      preDel = head
    } else if preDel != nil {
      preDel = preDel.Next
    }
  }
  
  if preDel == nil {
    if head.Next == nil {
      return nil
    } else {
      return head.Next
    }
  }
  
  preDel.Next = preDel.Next.Next
  return head
}