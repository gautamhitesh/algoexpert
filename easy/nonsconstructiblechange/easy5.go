//given array of coins find min amount of change you cannot create

package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{5, 7, 1, 1, 2, 3, 22}

	//solution brute force:

	// a sequence should run from 1 to n
	// sort the values
	//if array size =0 then min change you cannot create is 1
	// arr size is 1 and item is 1: ans 2
	// [2]:: 1
	// [1,1] :: 3
	// [2,1] :: 4

	// individual coins + sum of coins
	// sort the array
	// search for individual coin
	// search for sum of coins 2,3,4 ...
	//[5,7,1,1,2,3,22]
	// coins: 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19, <20> , <21>, 22

	// [1,1,2,3,5,7,22]
	// [1,2,4,7,12,19,41]
	//map of coins
	min_change := 1
	sort.Ints(coins)
	fmt.Println(coins)
	for i := 0; i < len(coins); i++ {
		if coins[i] <= min_change {
			min_change += coins[i]
		} else {
			break
		}
	}
	fmt.Println(min_change)
}
