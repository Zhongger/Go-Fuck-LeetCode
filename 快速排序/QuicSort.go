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
	start1 := time.Now().UnixNano()/1e6
	QuickSortX(arr1, 0, len(arr1)-1)
	end1 := time.Now().UnixNano()/1e6
	fmt.Printf("10万随机数的排序耗时为：%d \n", end1-start1)

	//20万随机数
	arr2 := make([]int, 200000)
	for i := 0; i < 200000; i++ {
		result := rand2.Intn(200000)
		arr2 = append(arr2, result)
	}
	start2 := time.Now().UnixNano()/1e6
	QuickSortX(arr2, 0, len(arr2)-1)
	end2 := time.Now().UnixNano()/1e6
	fmt.Printf("20万随机数的排序耗时为：%d \n", end2-start2)

	//40万随机数
	arr4 := make([]int, 400000)
	for i := 0; i < 400000; i++ {
		result := rand2.Intn(400000)
		arr4 = append(arr4, result)
	}
	start4 := time.Now().UnixNano()/1e6
	QuickSortX(arr4, 0, len(arr4)-1)
	end4 := time.Now().UnixNano()/1e6
	fmt.Printf("40万随机数的排序耗时为：%d \n", end4-start4)

	//70万随机数
	arr7 := make([]int, 700000)
	for i := 0; i < 700000; i++ {
		result := rand2.Intn(700000)
		arr7 = append(arr7, result)
	}
	start7 := time.Now().UnixNano()/1e6
	QuickSortX(arr7, 0, len(arr7)-1)
	end7 := time.Now().UnixNano()/1e6
	fmt.Printf("70万随机数的排序耗时为：%d \n", end7-start7)

	//100万随机数
	arr10 := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		result := rand2.Intn(1000000)
		arr10 = append(arr10, result)
	}
	start10 := time.Now().UnixNano()/1e6
	QuickSortX(arr10, 0, len(arr10)-1)
	end10 := time.Now().UnixNano()/1e6
	fmt.Printf("100万随机数的排序耗时为：%d \n", end10-start10)

	//200万随机数
	arr20 := make([]int, 2000000)
	for i := 0; i < 2000000; i++ {
		result := rand2.Intn(2000000)
		arr20 = append(arr20, result)
	}
	start20 := time.Now().UnixNano()/1e6
	QuickSortX(arr20, 0, len(arr20)-1)
	end20 := time.Now().UnixNano()/1e6
	fmt.Printf("200万随机数的排序耗时为：%d \n", end20-start20)

	//400万随机数
	arr40 := make([]int, 4000000)
	for i := 0; i < 4000000; i++ {
		result := rand2.Intn(4000000)
		arr40 = append(arr40, result)
	}
	start40 := time.Now().UnixNano()/1e6
	QuickSortX(arr40, 0, len(arr40)-1)
	end40 := time.Now().UnixNano()/1e6
	fmt.Printf("400万随机数的排序耗时为：%d \n", end40-start40)

	//600万随机数
	arr60 := make([]int, 6000000)
	for i := 0; i < 6000000; i++ {
		result := rand2.Intn(6000000)
		arr60 = append(arr60, result)
	}
	start60 := time.Now().UnixNano()/1e6
	QuickSortX(arr60, 0, len(arr60)-1)
	end60 := time.Now().UnixNano()/1e6
	fmt.Printf("600万随机数的排序耗时为：%d \n", end60-start60)

	//900万随机数
	arr90 := make([]int, 9000000)
	for i := 0; i < 9000000; i++ {
		result := rand2.Intn(9000000)
		arr90 = append(arr90, result)
	}
	start90 := time.Now().UnixNano()/1e6
	QuickSortX(arr90, 0, len(arr90)-1)
	end90 := time.Now().UnixNano()/1e6
	fmt.Printf("900万随机数的排序耗时为：%d \n", end90-start90)
}

// 如果数量小于13直接用插入排序
func SortForMerge(arr []int, left, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i]
		var j int
		for j = i; j > left && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func swap(arr []int, i, j int) { // 数据交换
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSortX(arr []int, left, right int) {
	if right-left < 2 {
		SortForMerge(arr, left, right) // 如果数量小于3直接用插入排序
	} else { // 使用快速排序
		// 随机找一个数字

		swap(arr, left, rand2.Int()%(right-left+1)+left)
		vdata := arr[left] // 坐标数据，比我小往左，大往右
		lt := left         // < vdata
		gt := right + 1    // > vdata
		i := left + 1      // == vdata
		for i < gt {
			if arr[i] > vdata {
				swap(arr, gt-1, i) //移动到大于的地方
				gt--
			} else if arr[i] < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++
			} else {
				i++
			}
		}
		swap(arr, lt, left)
		QuickSortX(arr, left, lt-1) //递归处理小于那一段
		QuickSortX(arr, gt, right)  //递归处理大于那一段
	}
}
