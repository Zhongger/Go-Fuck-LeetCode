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
	QuickSort(arr1, 0, len(arr1)-1)
	end1 := time.Now().Unix()
	fmt.Printf("10万随机数的排序耗时为：%d \n", end1-start1)

	//20万随机数
	arr2 := make([]int, 200000)
	for i := 0; i < 200000; i++ {
		result := rand2.Intn(200000)
		arr2 = append(arr2, result)
	}
	start2 := time.Now().Unix()
	QuickSort(arr2, 0, len(arr2)-1)
	end2 := time.Now().Unix()
	fmt.Printf("20万随机数的排序耗时为：%d \n", end2-start2)

	//40万随机数
	arr4 := make([]int, 400000)
	for i := 0; i < 400000; i++ {
		result := rand2.Intn(400000)
		arr4 = append(arr4, result)
	}
	start4 := time.Now().Unix()
	QuickSort(arr4, 0, len(arr4)-1)
	end4 := time.Now().Unix()
	fmt.Printf("40万随机数的排序耗时为：%d \n", end4-start4)

	//70万随机数
	arr7 := make([]int, 700000)
	for i := 0; i < 700000; i++ {
		result := rand2.Intn(700000)
		arr7 = append(arr7, result)
	}
	start7 := time.Now().Unix()
	QuickSort(arr7, 0, len(arr7)-1)
	end7 := time.Now().Unix()
	fmt.Printf("70万随机数的排序耗时为：%d \n", end7-start7)

	//100万随机数
	arr10 := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		result := rand2.Intn(1000000)
		arr10 = append(arr10, result)
	}
	start10 := time.Now().Unix()
	QuickSort(arr10, 0, len(arr10)-1)
	end10 := time.Now().Unix()
	fmt.Printf("100万随机数的排序耗时为：%d \n", end10-start10)

	//200万随机数
	arr20 := make([]int, 2000000)
	for i := 0; i < 2000000; i++ {
		result := rand2.Intn(2000000)
		arr20 = append(arr20, result)
	}
	start20 := time.Now().Unix()
	QuickSort(arr20, 0, len(arr20)-1)
	end20 := time.Now().Unix()
	fmt.Printf("200万随机数的排序耗时为：%d \n", end20-start20)

	//400万随机数
	arr40 := make([]int, 4000000)
	for i := 0; i < 4000000; i++ {
		result := rand2.Intn(4000000)
		arr40 = append(arr40, result)
	}
	start40 := time.Now().Unix()
	QuickSort(arr40, 0, len(arr40)-1)
	end40 := time.Now().Unix()
	fmt.Printf("400万随机数的排序耗时为：%d \n", end40-start40)

	//600万随机数
	arr60 := make([]int, 6000000)
	for i := 0; i < 6000000; i++ {
		result := rand2.Intn(6000000)
		arr60 = append(arr60, result)
	}
	start60 := time.Now().Unix()
	QuickSort(arr60, 0, len(arr60)-1)
	end60 := time.Now().Unix()
	fmt.Printf("600万随机数的排序耗时为：%d \n", end60-start60)

	//900万随机数
	arr90 := make([]int, 9000000)
	for i := 0; i < 9000000; i++ {
		result := rand2.Intn(9000000)
		arr90 = append(arr90, result)
	}
	start90 := time.Now().Unix()
	QuickSort(arr90, 0, len(arr90)-1)
	end90 := time.Now().Unix()
	fmt.Printf("900万随机数的排序耗时为：%d \n", end90-start90)
}

func QuickSort(values []int, left int, right int) {
	if  left < right {

		// 设置基准值
		temp := values[left]

		// 设置哨兵
		i, j := left, right
		for  {
			// 从右向左找，找到第一个比基准值小的数
			for  values[j] >= temp && i < j {
				j--
			}

			// 从左向右找，找到第一个比基准值大的数
			for  values[i] <= temp && i < j {
				i++
			}

			// 如果哨兵相遇，则退出循环
			if  i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}

		// 将基准值移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		QuickSort(values, left, i-1)
		QuickSort(values, i+1, right)
	}
}