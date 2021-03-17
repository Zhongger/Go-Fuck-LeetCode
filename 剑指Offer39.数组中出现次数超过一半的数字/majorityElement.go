package main

func main() {
	//nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	nums := []int{1}
	println(majorityElement(nums))
}
func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	length := len(nums) / 2
	for _, num := range nums {
		if numMap[num] >= length {
			return num
		} else {
			numMap[num] ++
		}
	}
	return -1
}
