/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

    first := &ListNode{}
    second := &ListNode{}

    s_half := second
    f_half := first
    current := head

    for current != nil {
        if current.Val < x {
            first.Next = current
            first = first.Next
        } else {
            second.Next = current
            second = second.Next
        }
        current = current.Next
    }
    second.Next = nil

    first.Next = s_half.Next

    return f_half.Next
}