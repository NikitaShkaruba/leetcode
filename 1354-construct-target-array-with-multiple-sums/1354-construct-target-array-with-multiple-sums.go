type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func isPossible(target []int) bool {
	maxHeap, sum := new(MaxHeap), 0
	for _, n := range target {
		heap.Push(maxHeap, n)
		sum += n
	}

	for {
		a := heap.Pop(maxHeap).(int)
		sum -= a
    
		if a == 1 || sum == 1 {
			return true
		}
		if sum >= a || sum == 0 || a%sum == 0 {
			return false
		}
    
		a %= sum
		sum += a
    heap.Push(maxHeap, a)
	}
}