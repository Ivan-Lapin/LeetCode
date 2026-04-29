type MinStack struct {
	head *Node
}

type Node struct {
	Val  int
	Min  int
	Next *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := val

	if this.head != nil && this.head.Min < minVal {
		minVal = this.head.Min
	}

	node := &Node{
		Val:  val,
		Min:  minVal,
		Next: this.head,
	}

	this.head = node
}

func (this *MinStack) Pop() {
    if this.head != nil {
        this.head = this.head.Next
    }
}

func (this *MinStack) Top() int {
    if this.head == nil {
        return 0
    }
	return this.head.Val
}

func (this *MinStack) GetMin() int {
    if this.head == nil {
        return 0
    }
	return this.head.Min
}