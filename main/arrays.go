package main

import (
	"fmt"
	"sort"
)

var Nums [5]int

func Start() {
	// fmt.Println(Nums)
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 1, 88))
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 4, 77))
	// fmt.Println(SafeWrite([5]int{0, 1, 2, 3, 4}, 5, 77))
	// fmt.Println(SafeWrite([5]int{}, 5, 10))
	fmt.Println(Remove(NumsSlice, 1))
	fmt.Println(UniqueSortedUserIDs([]int64{1, 2, 2, 4, 5, 4, 7}))
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

func IntsCopy(src []int, maxLen int) []int {
	if maxLen == 0 || maxLen < 0 {
		return []int{}
	} else if maxLen > len(src) {
		srcCopy := make([]int, len(src))
		copy(srcCopy, src)
		return srcCopy
	} else {
		srcCopy := make([]int, maxLen)
		copy(srcCopy, src)
		return srcCopy
	}
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) == 0 {
		return []int64{}
	}
	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})
	prev := 1
	for curr := 1; curr < len(userIDs); curr++ {
		if userIDs[curr-1] != userIDs[curr] {
			userIDs[prev] = userIDs[curr]
			prev++
		}
	}
	return userIDs[:prev]
}
