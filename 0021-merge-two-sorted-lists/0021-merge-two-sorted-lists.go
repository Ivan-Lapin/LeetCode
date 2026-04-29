/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func recursionForMergeTwoLists(node1 *ListNode, node2 *ListNode, list *ListNode) {
	if node1 == nil && node2 == nil {
		return
	}

	if node1 == nil  {
		list.Next = node2
		node2 = node2.Next
		recursionForMergeTwoLists(node1, node2, list.Next)
    
	} else if node2 == nil {
		list.Next = node1
		node1 = node1.Next
		recursionForMergeTwoLists(node1, node2, list.Next)

	} else if node1.Val < node2.Val {
        list.Next = node1
		node1 = node1.Next
		recursionForMergeTwoLists(node1, node2, list.Next)

    } else {
        list.Next = node2
		node2 = node2.Next
		recursionForMergeTwoLists(node1, node2, list.Next)
    }

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}

	recursionForMergeTwoLists(list1, list2, list)

	return list.Next
}