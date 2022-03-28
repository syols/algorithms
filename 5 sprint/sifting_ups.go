package main

type Heap struct {
	heap []int
}

const StartIndex = 1


func siftUp(heap []int, idx int) int {
	h := Heap{
		heap: heap,
	}
	return h.up(idx)
}


func (h *Heap) up(index int) int{
	pos := index
	for index != StartIndex {
		parentIndex := index / 2
		if h.less(parentIndex, index) {
			pos = parentIndex
			h.swap(parentIndex, index)
		}
		index = parentIndex
	}
	return pos
}

func (h *Heap) swap(left int, right int) {
	h.heap[right], h.heap[left] = h.heap[left], h.heap[right]
}

func (h *Heap) less(parentIndex int, childIndex int) bool {
	first := h.heap[parentIndex]
	second := h.heap[childIndex]
	return first < second
}
//
//func main() {
//	sample := []int{-1, 12, 6, 8, 3, 15, 7}
//	pos := siftUp(sample, 5)
//	if pos != 1 {
//		panic("WA")
//	}
//
//	pos = siftUp(sample, 3)
//	if  pos != 3 {
//		panic("WA")
//	}
//}