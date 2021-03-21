package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Printf("%v", res)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	var res []int
	for i, num := range nums {
		numMap[num] = i
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		otherNum := target - num
		index, isExit := numMap[otherNum]
		if isExit {
			res = append(res, num, nums[index])
			break
		}
	}
	return res

}
