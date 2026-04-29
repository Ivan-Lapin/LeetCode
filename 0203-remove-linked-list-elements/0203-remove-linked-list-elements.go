/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    
    dummy := &ListNode{Next: head}
    curr := dummy

    for curr != nil {
        next := curr.Next
        for next != nil && next.Val == val {
            next_next := next.Next
            next = next_next
        }
        curr.Next = next
        curr = curr.Next
    }

    return dummy.Next
}