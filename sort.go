package main

import "fmt"

//冒泡排序
func bubbleSort(arr []int) {
	lenArr := len(arr)
	for i := 0; i < lenArr-1; i++ {
		for j := 0; j < lenArr-1-i; j++ {
			// 相邻元素两两对比
			if arr[j] > arr[j+1] {
				// 元素交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//选择排序
func selectSort(arr []int) {
	lenArr := len(arr)
	minIndex := 0
	for i := 0; i < lenArr-1; i++ {
		minIndex = i
		for j := i + 1; j < lenArr; j++ {
			// 寻找最小的数
			if arr[j] < arr[minIndex] {
				// 将最小数的索引保存
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func insertSort(arr []int) {
	lenArr := len(arr)
	preIndex, current := 0, 0
	for i := 1; i < lenArr; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}

func partitionArr(arr []int, left int, right int) int {
	// arr 待分割序列，左边界，右边界，分割后轴位置
	temp := arr[right]
	for left != right {
		for arr[left] <= temp && left < right {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			// 赋值后，left不动，right左移
			right--
		}
		for arr[right] >= temp && right > left {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			// 赋值后 right不懂，left右移
			left++
		}
	}
	// 当left == right，把轴值放回left上
	arr[left] = temp
	return left
}

func quickSort(arr []int, left int, right int) {
	//子序列剩下0 或 1个元素，排序结束
	if right <= left {
		return
	}
	// 选择数组中间作为轴
	var pivot int
	pivot = (left + right) / 2
	fmt.Println(pivot)
	// 把轴值放在数组最后
	arr[right], arr[pivot] = arr[pivot], arr[right]
	fmt.Println(arr)
	// 分割后，左边的值 < 轴值 < 右边值
	pivot = partitionArr(arr, left, right)
	fmt.Println(pivot, left, right, arr)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func TopK(arr []int, top int) int {
	return quickSelect(arr, 2, 0, len(arr)-1)
}
func quickSelect(arr []int, top, left, right int) int {
	if left == right {
		return arr[right]
	}
	index := partitionArr(arr, left, right)
	if index-left+1 > top {
		return quickSelect(arr, top, left, index-1)
	} else if index-left+1 == top {
		return arr[index]
	} else {
		return quickSelect(arr, top-index+left-1, index+1, right)
	}

}

func main() {
	arr := []int{22, 33, 49, 47, 33, 12, 68, 29}
	//bubbleSort(arr)
	//selectSort(arr)
	//insertSort(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr2 := []int{22, 33, 49, 47, 33, 12, 68, 29}
	fmt.Println(TopK(arr2, 2))

}
