type MyStack struct {
    queue []int
}


func Constructor() MyStack {
    return MyStack{
        queue: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
    num := this.queue[len(this.queue)-1]
    this.queue = this.queue[:len(this.queue)-1]
    return num
}


func (this *MyStack) Top() int {
    return this.queue[len(this.queue)-1]
}


func (this *MyStack) Empty() bool {
    if len(this.queue) == 0 {
        return true
    }
    return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */