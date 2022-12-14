package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{6, 5, 3, -1, 8, 7, -2, 4, 4}
	fmt.Println("[UNSORTED]:\t", arr)
	fmt.Println("[SORTED]:\t", mergeSort(arr))
}

// 二路归并排序算法
func mergeSort(arr []int) []int {
	// 递归结束条件
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	arr = merge(left, right)
	fmt.Println("[DEBUG merged]:\t", arr)
	return arr
}

func merge(left, right []int) []int {
	// 遍历比较后合并两个数组
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		// 数组提前排序完毕
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		// 比较更小的值追加到 res[] 中
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}
