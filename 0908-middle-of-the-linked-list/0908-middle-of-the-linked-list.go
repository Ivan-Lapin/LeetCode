/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slice := []*ListNode{head}
    return findMiddel(head.Next, slice)
}

func findMiddel(head *ListNode, slice []*ListNode) *ListNode {
    if head == nil {
        return slice[len(slice)/2]
    }
    slice = append(slice, head)
    return findMiddel(head.Next, slice)
}