func reorderList(head *ListNode)  {

    if head == nil || head.Next == nil {
		return
	}

    slow, fast := head, head

    for fast.Next != nil && fast.Next.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }

    current := slow.Next
    slow.Next = nil

    var prev *ListNode

    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }

    start := head
    second := prev

    for second != nil {
		next1 := start.Next
		next2 := second.Next

		start.Next = second
		second.Next = next1

		start = next1
		second = next2
	}

}