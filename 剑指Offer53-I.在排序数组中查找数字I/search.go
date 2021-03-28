package main

import "fmt"

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	res := search(nums, 5)
	fmt.Printf("%v", res)
}

func search(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	var res = 0

	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return res

}
