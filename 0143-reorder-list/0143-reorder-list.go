/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {

    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    second := slow.Next
    slow.Next = nil

    var prev *ListNode
    for second != nil {
        next := second.Next
        second.Next = prev
        prev = second
        second = next
    }

    first, second := head, prev

    for second != nil {

        tmp1, tmp2 := first.Next, second.Next

        first.Next = second
        second.Next = tmp1

        first, second = tmp1, tmp2

    }
}