package linkedList

type ListNode struct {
	val  interface{}
	next *ListNode
}

// with tail pointer
type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewListNode(val interface{}) *ListNode {
	return &ListNode{val: val}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// returns the number of data elements in the list
func (l *LinkedList) Size() int {
	return l.size
}

// returns true if empty
func (l *LinkedList) Empty() bool {
	return l.size == 0
}

// returns the value of the nth item (starting at 0 for first)
func (l *LinkedList) ValueAt(index int) interface{} {
	node := l.head
	for node != nil {
		if index == 0 {
			return node.val
		}
		node = node.next
		index--
	}
	return nil
}

// adds an item to the front of the list
func (l *LinkedList) PushFront(val interface{}) {
	node := NewListNode(val)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.size++
}

// remove the front item and return its value
func (l *LinkedList) PopFront() interface{} {
	if l.head == nil {
		return nil
	}
	val := l.head.val
	l.head = l.head.next
	l.size--
	return val
}

// adds an item at the end
func (l *LinkedList) PushBack(val interface{}) {
	node := NewListNode(val)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

// removes end item and returns its value
func (l *LinkedList) PopBack() interface{} {
	if l.head == nil {
		return nil
	}
	node := l.head
	var prev *ListNode
	for node.next != nil {
		prev = node
		node = node.next
	}
	if prev != nil {
		prev.next = nil
		l.tail = prev
	} else {
		l.head = nil
		l.tail = nil
	}
	l.size--
	return node.val
}

// get the value of the front item
func (l *LinkedList) Front() interface{} {
	if l.head == nil {
		return nil
	}
	return l.head.val
}

// get the value of the end item
func (l *LinkedList) Back() interface{} {
	if l.head == nil {
		return nil
	}
	return l.tail.val
}

// insert value at index, so the current item at that index is pointed to by the new item at the index
func (l *LinkedList) Insert(index int, val interface{}) {
	if index == 0 {
		l.PushFront(val)
		return
	}
	node := l.head
	for node != nil {
		if index == 1 {
			newNode := NewListNode(val)
			newNode.next = node.next
			node.next = newNode
			l.size++
			return
		}
		node = node.next
		index--
	}
	panic("Index out of range")
}

// removes node at given index
func (l *LinkedList) Erase(index int) {
	if index == 0 {
		l.PopFront()
		return
	}
	node := l.head
	var prev *ListNode
	for node != nil {
		if index == 0 {
			prev.next = node.next
			l.size--
			return
		}
		prev = node
		node = node.next
		index--
	}
	panic("Index out of range")
}

// returns the value of the node at the nth position from the end of the list
func (l *LinkedList) ValueNFromEnd(n int) interface{} {
	// node := l.head
	index := l.size - n - 1
	return l.ValueAt(index)
	// for i := 0; i < l.size-n-1; i++ {
	// 	node = node.next
	// }
	// return node.val
}

// reverses the list
func (l *LinkedList) Reverse() {
	var prev *ListNode
	node := l.head
	for node != nil {
		next := node.next
		node.next = prev
		prev = node
		node = next
	}
	l.head = prev
}

// removes the first item in the list with this value
func (l *LinkedList) RemoveValue(val interface{}) {
	node := l.head
	var prev *ListNode
	for node != nil {
		if node.val == val {
			if prev != nil {
				prev.next = node.next
			} else {
				l.head = node.next
			}
			l.size--
			return
		}
		prev = node
		node = node.next
	}
}
