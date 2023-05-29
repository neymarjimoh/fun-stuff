package minHeap

type MinHeap struct {
	data []int
}

type IMinHeap interface {
	Poll() int
	Push(n int)
	Peek() int
	Size() int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
	}
}

func (h *MinHeap) Push(n int) {
	h.data = append(h.data, n)
	h.heapifyUp()
}

func (h *MinHeap) heapifyUp() {
	childIndex := h.Size() - 1
	parentIndex := h.getParentIndex(childIndex)
	for h.data[childIndex] < h.data[parentIndex] {
		h.data[childIndex], h.data[parentIndex] = h.data[parentIndex], h.data[childIndex]
		childIndex = parentIndex
		parentIndex = h.getParentIndex(childIndex)
		if parentIndex < 0 {
			break
		}
	}
}

func (h *MinHeap) Poll() (int, bool) {
	var rootNode int
	if h.Size() == 0 {
		return rootNode, false
	}
	rootNode = h.data[0]
	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	h.heapifyDown()
	return rootNode, true
}

func (h *MinHeap) heapifyDown() {
	currIndex := 0
	leftIndex, rightIndex := h.getLeftChildIndex(currIndex), h.getRightChildIndex(currIndex)
	indexToSwap := leftIndex

	for leftIndex < h.Size() {
		if rightIndex < h.Size() && h.data[rightIndex] < h.data[leftIndex] {
			indexToSwap = rightIndex
		} else {
			indexToSwap = leftIndex
		}
		if h.data[currIndex] <= h.data[indexToSwap] {
			break
		}
		h.data[currIndex], h.data[indexToSwap] = h.data[indexToSwap], h.data[currIndex]
		currIndex = indexToSwap
		leftIndex = h.getLeftChildIndex(currIndex)
	}
}

func (h *MinHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MinHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MinHeap) Peek() int {
	return h.data[0]
}

func (h *MinHeap) Size() int {
	return len(h.data)
}
