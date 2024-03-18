package main

import "fmt"

// See task readme
func main() {

	q := []int{2, 0, 1}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
	q = []int{1, 2, 3}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
	q = []int{0, 1, 2}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
	q = []int{2, 1, 0}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
	q = []int{4, 1, 0, 1, 4}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
	q = []int{4, 1, 0, 4, 1}
	fmt.Printf("q - [% v]; ice cream min number - %v\n", q, numberOfIceCreams(q))
}

func numberOfIceCreams(queue []int) int {
	num := 0
	// stores number of ice creams for each person in queue
	nums := make([]int, len(queue))

	// not include last item, because we look current and next items
	for i := 0; i < len(queue)-1; i++ {
		if queue[i] < queue[i+1] {
			updateFront(nums[:i+1], -1)
		} else if queue[i] > queue[i+1] {
			updateFront(nums[:i+1], 1)
		}
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	if min <= 0 {
		for i := 0; i < len(nums); i++ {
			num += (nums[i] - min) + 1
		}
	} else {
		for _, n := range nums {
			num += n
		}
	}
	return num
}

func updateFront(queue []int, n int) {
	if !(n == 1 || n == -1) {
		panic("updaet fron n is 1 or -1 only")
	}
	for i := len(queue) - 1; i >= 0; i-- {
		queue[i] += n
	}
}
