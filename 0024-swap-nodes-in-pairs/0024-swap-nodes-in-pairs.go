/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    old := &ListNode{Next: head}
    prev := old

    for prev.Next != nil && prev.Next.Next != nil {
        first := prev.Next
        second := first.Next

        first.Next = second.Next
        second.Next = first

        prev.Next = second 
        prev = first
    }

    return old.Next
}