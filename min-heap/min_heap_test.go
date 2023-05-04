package minHeap

import (
	"reflect"
	"testing"
)

func TestPushMinHeap(t *testing.T) {
	testCases := map[string]struct {
		input []int
		expected []int
	}{
		"given an array return the heapified array": {
			[]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7},
			[]int{4, 6, 5, 8, 7, 10, 13, 12, 11, 9},
		},
		"given an empty array return an empty array": {
			[]int{},
			[]int{},
		},
		"given an array with one element, return same array": {
			[]int{3},
			[]int{3},
		},
	}
	
	for name, tc := range testCases {
		var heap = NewMinHeap();

		t.Run(name, func(t *testing.T) {
			t.Parallel();
			for _, v := range tc.input {
				heap.Push(v);
			}
			if !reflect.DeepEqual(heap.data, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, heap.data);
			}
		});
	}
}

func TestPollMinHeap(t *testing.T) {
	testCases := map[string]struct {
		input []int
	}{
		"given an array poll out root node": {
			[]int{6, 5, 4, 8, 9, 10, 13, 12, 11, 7},
		},
		"given an empty array return an empty array": {
			[]int{},
		},
		"given an array with one element, poll out the element": {
			[]int{3},
		},
	}

	for name, tc := range testCases {
		var heap = NewMinHeap();

		t.Run(name, func(t *testing.T) {
			t.Parallel();
			for _, v := range tc.input {
				heap.Push(v);
			}
			for _, v := range heap.data {
				if val, ok := heap.Poll(); val != v && ok {
					t.Errorf("expected %v, got %v", val, v);
				}
			}
		});
	}
}
