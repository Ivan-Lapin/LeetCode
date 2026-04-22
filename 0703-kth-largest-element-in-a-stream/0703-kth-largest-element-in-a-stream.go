import "container/heap"

type MinHeap []int

func (heap MinHeap) Len() int {
    return len(heap)
}

func (heap MinHeap) Less(i, j int) bool {
    return heap[i] < heap[j]
}

func (heap MinHeap) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MinHeap) Push(x interface{}) {
    *heap = append(*heap, x.(int))
}

func (heap *MinHeap) Pop() interface{} {
    old := *heap
    x := old[len(old)-1]
    *heap = old[:len(old)-1]
    return x
}   

type KthLargest struct {
    k int
    heap MinHeap
}


func Constructor(k int, nums []int) KthLargest {

    h := MinHeap{}
    heap.Init(&h)
    kth := KthLargest{
        k: k,
        heap: h,
    }

    for _, num := range nums {
        kth.Add(num)
    }

    return kth
}


func (this *KthLargest) Add(val int) int {
    if this.heap.Len() < this.k {
        heap.Push(&this.heap, val)
    } else if val > this.heap[0]{
        heap.Pop(&this.heap)
        heap.Push(&this.heap, val)
    }

    return this.heap[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */