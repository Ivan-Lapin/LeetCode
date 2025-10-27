
type MaxHeap []int

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i int, j int) bool {return h[i] > h[j]}
func (h MaxHeap) Swap(i int, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x any) {*h = append(*h, x.(int))}
func (h *MaxHeap) Pop() any {
    old := *h
    n := len(*h)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func lastStoneWeight(stones []int) int {
    h := MaxHeap(stones)
    heap.Init(&h)

    for h.Len() > 1 {
        y := heap.Pop(&h).(int)
        x := heap.Pop(&h).(int)
        if y != x {
            heap.Push(&h, (y-x))
        }
    }

    if h.Len() == 0 {
        return 0
    }

    return heap.Pop(&h).(int)
    
}