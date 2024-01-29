package array

import "testing"

func TestArrayGetSize(t *testing.T) {
	a := NewArray()
	if size := a.GetSize(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	a.Push(1)
	if size := a.GetSize(); size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}
}

func TestArrayIsEmpty(t *testing.T) {
	a := NewArray()
	if !a.IsEmpty() {
		t.Error("Expected array to be empty")
	}

	a.Push(1)
	if a.IsEmpty() {
		t.Error("Expected array not to be empty")
	}
}

func TestArrayGetCapacity(t *testing.T) {
	a := NewArray(WithCapacity(3))
	if capacity := a.GetCapacity(); capacity != 3 {
		t.Errorf("Expected capacity 3, got %d", capacity)
	}
}

func TestArrayAt(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)

	if val := a.At(0); val != 1 {
		t.Errorf("Expected value 1 at index 0, got %d", val)
	}

	if val := a.At(1); val != 2 {
		t.Errorf("Expected value 2 at index 1, got %d", val)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds index, but it did not panic")
		}
	}()

	a.At(2)
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestArrayInsert(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)

	a.Insert(1, 3)

	expected := []int{1, 3, 2}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after insertion, got %v", expected, a.data)
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds index in Insert, but it did not panic")
		}
	}()

	a.Insert(5, 4)
}

func TestArrayResize(t *testing.T) {
	a := NewArray(WithCapacity(3))
	a.Push(1)
	a.Push(2)

	a.resize(6)

	if capacity := a.GetCapacity(); capacity != 6 {
		t.Errorf("Expected capacity 6 after resize, got %d", capacity)
	}

	if size := a.GetSize(); size != 2 {
		t.Errorf("Expected size 2 after resize, got %d", size)
	}
}

func TestArrayPush(t *testing.T) {
	a := NewArray(WithCapacity(2))
	a.Push(1)
	a.Push(2)
	a.Push(3)

	expected := []int{1, 2, 3}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after pushing elements, got %v", expected, a.data)
	}
}

func TestArrayFind(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)
	a.Push(3)

	if index := a.Find(2); index != 1 {
		t.Errorf("Expected index 1 for value 2, got %d", index)
	}

	if index := a.Find(4); index != -1 {
		t.Errorf("Expected index -1 for value 4 (not found), got %d", index)
	}
}

func TestArrayRemove(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)
	a.Push(3)

	a.Remove(2)

	expected := []int{1, 3}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after removing element, got %v", expected, a.data)
	}

	// Test removing non-existing element
	a.Remove(5)

	if size := a.GetSize(); size != 2 {
		t.Errorf("Expected size 2 after attempting to remove non-existing element, got %d", size)
	}
}

func TestArrayDelete(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)
	a.Push(3)

	a.Delete(1)

	expected := []int{1, 3}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after deleting element, got %v", expected, a.data)
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds index in Delete, but it did not panic")
		}
	}()

	a.Delete(5)
}

func TestArrayPrepend(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)

	a.Prepend(3)

	expected := []int{3, 1, 2}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after prepending element, got %v", expected, a.data)
	}
}

func TestArrayPop(t *testing.T) {
	a := NewArray()
	a.Push(1)
	a.Push(2)

	popped := a.Pop()

	if popped != 2 {
		t.Errorf("Expected popped value 2, got %d", popped)
	}

	expected := []int{1}
	if !equalSlices(a.data, expected) {
		t.Errorf("Expected %v after popping element, got %v", expected, a.data)
	}

	// Test pop on an empty array
	emptyArray := NewArray()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for popping from an empty array, but it did not panic")
		}
	}()

	emptyArray.Pop()
}