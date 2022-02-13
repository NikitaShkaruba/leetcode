/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
  Test cases:
  
  Testing grounds:
  
    listA = [4,1]
                  [8, 4, 5]
    listB = [5,6,1]
    
  Solutions:
    1. make negative, iterate till negative, make negative: O(n) time, O(1) space
  
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
  negativateList(headA)
  
  var intersectNode *ListNode
  curB := headB
  for curB != nil && intersectNode == nil {
    if curB.Val < 0 {
      intersectNode = curB
    }
    curB = curB.Next
  }
  
  negativateList(headA)
  
  return intersectNode
}

func negativateList(head *ListNode) {
  cur := head
  
  for cur != nil {
    cur.Val *= -1
    cur = cur.Next
  }
}