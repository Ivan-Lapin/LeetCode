/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

    if head == nil {
        return false
    }

    if head.Next == nil {
        return true
    }

    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    var second *ListNode

    if fast == nil {
        second = slow
        slow = nil
    } else {
        second = slow.Next
        slow.Next = nil 
    }

    var prev *ListNode
    for second != nil {
        next := second.Next
        second.Next = prev
        prev = second
        second = next
    }

    first := head
    for first != nil && prev != nil{
        if first.Val != prev.Val {
            return false
        }
        first = first.Next
        prev = prev.Next
    }

    return true
}