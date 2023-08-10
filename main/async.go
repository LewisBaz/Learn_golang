package main

import (
	"fmt"
	"time"
)

func StartAsync() {
	fmt.Println(MaxSum([]int{1, 2, 3}, []int{3, 4, 5, 6}))
	fmt.Println()

	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{500, 5, 10, 25}

	res := <-sumCh // 540
	fmt.Println(res)
}

func calcSum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sumParallel(nums []int, res *int) {
	*res = calcSum(nums)
}

func MaxSum(nums1, nums2 []int) []int {
	s1, s2 := 0, 0

	go sumParallel(nums1, &s1)
	go sumParallel(nums2, &s2)

	time.Sleep(100 * time.Millisecond)

	if s2 > s1 {
		return nums2
	} else {
		return nums1
	}
}

func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sumCh <- calcSum(nums)
	}
}
