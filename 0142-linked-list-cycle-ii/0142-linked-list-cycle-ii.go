/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	m := make(map[*ListNode]bool)

	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	m[slow] = true

	for fast != nil && fast.Next != nil {
		slow = fast.Next
		fast = fast.Next.Next

		if _, exist := m[slow]; exist {
			return slow
		}
		m[slow] = true

		if _, exist := m[fast]; exist {
			return fast
		}
		m[fast] = true
	}

	return nil

}