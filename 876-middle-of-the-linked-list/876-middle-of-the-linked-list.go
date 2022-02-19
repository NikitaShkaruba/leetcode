/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*

  Solutions:
    1. Brute force: O(n) time, O(1) space
  
  Test cases:
    - 1*: 1
    - 1 -> 2*: 2
    - 1 -> 2 -> 3* -> 4 -> 5: 3
    - 1 -> 2 -> 3 -> 4* -> 5 -> 6: 4
  
  Testing grounds:
    
*/

func middleNode(head *ListNode) *ListNode {
  nodesAmount := 0
  currentNode := head
  for currentNode != nil {
    nodesAmount++
    currentNode = currentNode.Next
  }
  
  var middleIndex int
  if nodesAmount % 2 == 0 {
    middleIndex = nodesAmount / 2
  } else {
    middleIndex = int(math.Floor(float64(nodesAmount / 2)))
  }
  
  currentIndex := 0
  currentNode = head
  for currentIndex != middleIndex {
    currentNode = currentNode.Next
    currentIndex++
  }
  
  return currentNode
}