package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {

	originalSlice := newRandomSlice(0, 10, 10)
	fmt.Println("1. ", originalSlice)

	tmp := sliceExample(originalSlice)
	fmt.Printf("2.  %v, %p -> %p\n", tmp, originalSlice, tmp)

	tmp = addElements(originalSlice, 55)
	fmt.Printf("3.  %v, %p -> %p\n", tmp, originalSlice, tmp)

	tmp = copySlice(originalSlice)
	fmt.Printf("4. %v, %p -> %p\n", tmp, originalSlice, tmp)

	tmp = removeElement(originalSlice, 9)
	fmt.Printf("5. %v, %p -> %p\n", tmp, originalSlice, tmp)
}

func newRandomSlice(from, to, count int) []int {
	result := make([]int, count)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result[i] = rand.Intn(to-from) + from
	}

	return result
}

func sliceExample(arr []int) []int {

	result := []int{}

	for _, a := range arr {
		if a%2 == 0 {
			result = append(result, a)
		}
	}

	return result
}

func addElements(arr []int, x int) []int {
	result := make([]int, 0, len(arr)+1)

	return append(append(result, arr...), x)
}

func copySlice(arr []int) []int {
	return slices.Clone(arr)
}

func removeElement(arr []int, i int) []int {
	if 0 > i || i >= len(arr) {
		return []int{}
	}

	result := make([]int, 0, len(arr)-1)

	return append(append(result, arr[:i]...), arr[i+1:]...)
}
