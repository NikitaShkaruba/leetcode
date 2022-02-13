/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
 4 -> 5 -> 1 -> 9
      ^
      
 4 -> 1 -> 9
      ^
 */
func deleteNode(node *ListNode) {
  *node = *(node.Next)
}