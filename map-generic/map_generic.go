package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomMap[T comparable, V string | int] map[T]V

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	result := MapValues([]int{1, 2, 3}, func(i int) int {
		return i * 2
	})
	fmt.Printf("result: %v\n", result)

	m := make(CustomMap[string, int], 7)
	m["id"] = 1
}
