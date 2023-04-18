package datastructures

// Heap is our heap struct
type Heap struct {
	data []int
}

// NewHeap creates and returns a new heap
func NewHeap() *Heap {
	return &Heap{
		data: make([]int, 0),
	}
}

// Insert a new value onto our heap
func (h *Heap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

// Remove an item from our heap
func (h *Heap) Remove() int {
	if len(h.data) == 0 {
		panic("Heap is empty")
	}

	root := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.bubbleDown(0)
	return root
}

// bubbleUp an index
func (h *Heap) bubbleUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

// bubbleDown an index
func (h *Heap) bubbleDown(index int) {
	for {
		left := index*2 + 1
		right := index*2 + 2
		smallest := index

		if left < len(h.data) && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < len(h.data) && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}
