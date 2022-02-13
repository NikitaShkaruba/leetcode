/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode MergeTwoLists(ListNode headOne, ListNode headTwo) {
        if (headOne == null) {
          return headTwo;
        }
        if (headTwo == null) {
          return headOne;
        }
        if (headOne == null && headTwo == null) {
          return null;
        }
      
        ListNode mergedListHead;
        if (headOne.val <= headTwo.val) {
            mergedListHead = headOne;
            headOne = headOne.next;
        } else {
            mergedListHead = headTwo;
            headTwo = headTwo.next;
        }

        ListNode mergedListCurrent = mergedListHead;
        while (headOne != null || headTwo != null) {
            if (headOne != null && (headTwo == null || headOne.val <= headTwo.val)) {
                mergedListCurrent.next = headOne;
                headOne = headOne.next;
            } else {
                mergedListCurrent.next = headTwo;
                headTwo = headTwo.next;
            }

            mergedListCurrent = mergedListCurrent.next;
        }

        return mergedListHead;
    }
}