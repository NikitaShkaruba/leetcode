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
    - 1 -> 2 : 1 -> 2
    - 1 -> 1 -> 1 -> 2 : 1 -> 2
    - 1 -> 1  : 1
    - 1 : 1
    - nil -> nil (!!!)
  
  Testing grounds:
    - 1 -> 2 : 1 -> 2
           c

*/
func deleteDuplicates(head *ListNode) *ListNode {
  cur := head
  
  for cur != nil && cur.Next != nil {
    if cur.Val != cur.Next.Val {
      cur = cur.Next
    } else {
      cur.Next = cur.Next.Next
    }
  }
  
  return head
}