package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(intersection(a, b))

}

func intersection(s1, s2 []int) (bool, []int) {
	result := make([]int, 0, 0)

	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}

	for _, v := range s1 {
		if slices.Contains(s2, v) {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
