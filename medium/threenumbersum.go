//distinct number in an array
//non empty
//target value
// return array fo array of triplets equalling to that sum
//array should be in ascending order and so should be the triplets

//[12,3,1,2,-6,5,-8,6] , 0
//ans: [[2,6,-8],[3,5,-8],[1,5,-6]]
//

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	sort.Ints(arr)
	fmt.Println(arr)
	left := 1
	right := len(arr) - 1
	res := make([][]int, 0)
	counter := 0
	//iterate through array, counter, left=counter+1, right=len(arr)-1
	//if sum(3)<target left++ left<right
	//if sum(3)>target right-- right>left
	//if sum(3)==target counter++ counter<len(arr)-2

	for {
		if counter == len(arr)-2 {
			break
		}
		if left >= right {
			counter++
			left = counter + 1
			right = len(arr) - 1
			continue
		}
		temp := make([]int, 0)
		sum := arr[left] + arr[counter] + arr[right]
		fmt.Println(sum, arr[left], arr[counter], arr[right])
		if sum < target && left < right {
			left++
		} else if sum > target && right > left {
			right--
		} else if target == sum {
			temp = append(temp, arr[counter], arr[left], arr[right])
			res = append(res, temp)
			left++
			right--
		} else {
			counter++
			right = len(arr) - 1
			left = counter + 1
		}
	}
	fmt.Println(res)
}
