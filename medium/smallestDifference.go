package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{-1, 5, 10, 20, 28, 3}
	arr2 := []int{26, 134, 135, 15, 17}
	sort.Ints(arr1)
	sort.Ints(arr2)
	res := minDifference(arr1, arr2)
	fmt.Println(res)
}

//sort the arrays
func minDifference(arr1 []int, arr2 []int) []int {
	arr1C := 0
	arr2C := 0
	minDiff := 10000
	index1 := 0
	index2 := 0
	for {
		if arr1C == len(arr1) || arr2C == len(arr2) {
			break
		}
		diff := absDiff(arr1[arr1C], arr2[arr2C])
		fmt.Println(diff, arr1[arr1C], arr2[arr2C])
		if diff < minDiff {
			minDiff = diff
			index1 = arr1C
			index2 = arr2C
		} else {
			if arr1[arr1C] > arr2[arr2C] {
				arr2C++
			} else {
				arr1C++
			}
		}
	}
	res := make([]int, 0)
	res = append(res, arr1[index1], arr2[index2])
	fmt.Println(minDiff)
	return res
}

func absDiff(num1 int, num2 int) int {
	diff := 0
	if num1 < 0 {
		diff = num2 - num1
	} else {
		diff = num1 - num2
	}
	if diff < 0 {
		return -diff
	}
	return diff
}
