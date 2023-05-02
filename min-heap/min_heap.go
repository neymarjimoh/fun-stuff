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
    };
}

func (h *MinHeap) Push(n int) {
    h.data = append(h.data, n);
    h.heapifyUp();
}

func (h *MinHeap) heapifyUp() {
    childIndex := h.Size() - 1;
    parentIndex := h.getParentIndex(childIndex);
    for h.data[childIndex] < h.data[parentIndex] {
        h.data[childIndex], h.data[parentIndex] = h.data[parentIndex], h.data[childIndex];
        childIndex = parentIndex;
        parentIndex = h.getParentIndex(childIndex);
        if parentIndex < 0 {
            break;
        }
    }
}

func (h *MinHeap) Poll() int {
    
    return -1;
}

func (h *MinHeap) getParentIndex(childIndex int) int {
    return (childIndex - 1) / 2;
}

func (h *MinHeap) Peek() int {

    return -1;
}

func  (h *MinHeap) Size() int {
    return len(h.data);
}
