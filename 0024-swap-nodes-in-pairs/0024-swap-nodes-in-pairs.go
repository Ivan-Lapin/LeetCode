/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
    current := dummy 

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next
		after := second.Next

		current.Next = second
		second.Next = first
		first.Next = after

		current = first
	}

	return dummy.Next
}