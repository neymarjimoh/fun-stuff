package array

const indexOutOfBoundsError = "index out of bounds"

type Array struct {
	data          []int
	size          int
	capacity      int
	isCapacitySet bool
}

type Option func(a *Array)

func WithCapacity(capacity int) Option {
	return func(a *Array) {
		a.capacity = capacity
		a.isCapacitySet = true
	}
}

// NewArray creates a new instance of the Array type with an initial capacity set to 1.
func NewArray(opts ...Option) *Array {
	a := &Array{
		data:          make([]int, 0),
		size:          0,
		isCapacitySet: false,
		// capacity: 16,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// return the number of items it can hold
func (a *Array) GetCapacity() int {
	return a.capacity
}

// returns the item at a given index, blows up if index out of bounds
func (a *Array) At(index int) int {
	if index < 0 || index >= a.GetSize() {
		panic(indexOutOfBoundsError)
	}

	return a.data[index]
}

func (a *Array) IsFull() bool {
	return a.size == a.capacity
}

func (a *Array) Push(item int) {
	a.Insert(a.GetSize(), item)
}

func (a *Array) Insert(index, item int) {
	if index < 0 || index > a.GetSize() {
		panic(indexOutOfBoundsError)
	}

	if a.IsFull() && !a.isCapacitySet {
		a.resize(a.capacity * 2)
	}

	a.data = append(a.data, 0)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = item
	a.size++
}

func (a *Array) Prepend(item int) {
	a.Insert(0, item)
}

func (a *Array) Pop() int {
	if a.IsEmpty() {
		panic("array is empty")
	}

	lastIndex := a.GetSize() - 1
	item := a.data[lastIndex]
	a.data = a.data[:lastIndex]
	a.size--
	return item
}

func (a *Array) Delete(index int) {
	if index < 0 || index >= a.GetSize() {
		panic(indexOutOfBoundsError)
	}

	copy(a.data[index:], a.data[index+1:])
	a.data = a.data[:a.size-1]
	a.size--
}

func (a *Array) Find(item int) int {
	for i, v := range a.data {
		if v == item {
			return i
		}
	}

	return -1
}

func (a *Array) Remove(item int) {
	index := a.Find(item)
	if index != -1 {
		a.Delete(index)
	}
}

func (a *Array) resize(newCapacity int) {
	d := make([]int, newCapacity)
	copy(d, a.data)
	a.data = d
	a.capacity = newCapacity
}
