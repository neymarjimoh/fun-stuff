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

}

func (h *MinHeap) Poll() int {

	return -1;
}


func (h *MinHeap) Peek() int {

	return -1;
}

func  (h *MinHeap) Size() int {
	return len(h.data);
}

