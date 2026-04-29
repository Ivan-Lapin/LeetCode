/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    current := head
    tail := 0

    for current != nil {
        tail++
        current = current.Next
    }

    if tail == n {
        return head.Next
    }

    steps := tail-n

    current = head
    var next_next *ListNode
    for i := range steps {
        if current.Next != nil {
            next_next = current.Next.Next
        } else {
            next_next = nil
        }
        if i != steps-1 {
            current = current.Next
        }
    }
    current.Next = next_next

    return head
}