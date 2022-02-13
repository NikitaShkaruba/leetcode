/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  var front *ListNode
  mid := head
  end := head
  
  for mid != nil {
    end = mid.Next
    mid.Next = front
    front = mid
    mid = end
  }
  
  return front
}