package main

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	println(findRepeatNumber(nums))
}
func findRepeatNumber(nums []int) int {
	numberMap := make(map[int]bool)
	for _, num := range nums {
		_, isExit := numberMap[num]
		if isExit {
			return num
		} else {
			numberMap[num] = true
		}
	}
	return -1
}
