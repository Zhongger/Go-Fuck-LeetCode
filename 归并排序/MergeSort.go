package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func main() {
	//10万随机数
	arr1 := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		result := rand2.Intn(100000)
		arr1 = append(arr1, result)
	}
	start1 := time.Now().Unix()
	mergeSort(arr1)

	end1 := time.Now().Unix()
	fmt.Printf("10万随机数的排序耗时为：%d \n", end1-start1)

	//20万随机数
	arr2 := make([]int, 200000)
	for i := 0; i < 200000; i++ {
		result := rand2.Intn(200000)
		arr2 = append(arr2, result)
	}
	start2 := time.Now().Unix()
	mergeSort(arr2)
	end2 := time.Now().Unix()
	fmt.Printf("20万随机数的排序耗时为：%d \n", end2-start2)

	//40万随机数
	arr4 := make([]int, 400000)
	for i := 0; i < 400000; i++ {
		result := rand2.Intn(400000)
		arr4 = append(arr4, result)
	}
	start4 := time.Now().Unix()
	mergeSort(arr4)
	end4 := time.Now().Unix()
	fmt.Printf("40万随机数的排序耗时为：%d \n", end4-start4)

	//70万随机数
	arr7 := make([]int, 700000)
	for i := 0; i < 700000; i++ {
		result := rand2.Intn(700000)
		arr7 = append(arr7, result)
	}
	start7 := time.Now().Unix()
	mergeSort(arr7)
	end7 := time.Now().Unix()
	fmt.Printf("70万随机数的排序耗时为：%d \n", end7-start7)

	//100万随机数
	arr10 := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		result := rand2.Intn(1000000)
		arr10 = append(arr10, result)
	}
	start10 := time.Now().Unix()
	mergeSort(arr10)
	end10 := time.Now().Unix()
	fmt.Printf("100万随机数的排序耗时为：%d \n", end10-start10)

	//200万随机数
	arr20 := make([]int, 2000000)
	for i := 0; i < 2000000; i++ {
		result := rand2.Intn(2000000)
		arr20 = append(arr20, result)
	}
	start20 := time.Now().Unix()
	mergeSort(arr20)
	end20 := time.Now().Unix()
	fmt.Printf("200万随机数的排序耗时为：%d \n", end20-start20)

	//400万随机数
	arr40 := make([]int, 4000000)
	for i := 0; i < 4000000; i++ {
		result := rand2.Intn(4000000)
		arr40 = append(arr40, result)
	}
	start40 := time.Now().Unix()
	mergeSort(arr40)
	end40 := time.Now().Unix()
	fmt.Printf("400万随机数的排序耗时为：%d \n", end40-start40)

	//600万随机数
	arr60 := make([]int, 6000000)
	for i := 0; i < 6000000; i++ {
		result := rand2.Intn(6000000)
		arr60 = append(arr60, result)
	}
	start60 := time.Now().Unix()
	mergeSort(arr60)
	end60 := time.Now().Unix()
	fmt.Printf("600万随机数的排序耗时为：%d \n", end60-start60)

	//900万随机数
	arr90 := make([]int, 9000000)
	for i := 0; i < 9000000; i++ {
		result := rand2.Intn(9000000)
		arr90 = append(arr90, result)
	}
	start90 := time.Now().Unix()
	mergeSort(arr90)
	end90 := time.Now().Unix()
	fmt.Printf("900万随机数的排序耗时为：%d \n", end90-start90)
}
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
