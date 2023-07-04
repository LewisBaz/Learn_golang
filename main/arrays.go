package main

import (
	"fmt"
)

var Nums [5]int

func Start() {
	// fmt.Println(Nums)
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 1, 88))
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 4, 77))
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 5, 77))
	// fmt.Println(SafeWrite([5]int{}, 5, 10))
	fmt.Println(Remove(NumsSlice, 1))
}

func SafeWrite(nums [5]int, i, val int) [5]int {
	if len(nums) > i && i >= 0 {
		nums[i] = val
		return nums
	} else {
		return nums
	}
}

var NumsSlice = []int{1, 2, 3, 4, 5}

func Remove(nums []int, i int) []int {
	if len(nums) > i && i >= 0 {
		nums[i] = nums[len(nums)-1]
		return nums[:len(nums)-1]
	} else {
		return nums
	}
}

func Map(strs []string, mapFunc func(s string) string) []string {
	newStrs := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		newStrs[i] = mapFunc(strs[i])
	}
	return newStrs
}
