//find the tournament winner from array of arrays(size 2) having home,away pair
// result array has 0 then the away team has won. If 1 the home team has won

// e.g
// [["HTML", "C#"], ["C#", "Python"],["Python", "HTML"]
// results: [0,0,1]
// first match: C# // 0 showing away
// Second match: Python
// third match: Python

// Hence winner is Python

// cases: least--> 1 match for 2 teams
// 3 teams: matches --> 3
// 4 teams: matches --> 1,2 | 2,3 | 3,4 | 1,3 | 1,4 | 2,4

// Number of matches --> nCr // not required

//Solution: make array of teams and then conditionally add it as per results
// making array of teams from the first array of arrays

// search and add to new array

// use map instead of all above

package main

import "fmt"

func main() {
	// teams := make([]string, 0)
	// wins := make([]int, 0)
	matches := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	results := []int{0, 0, 1}
	wins := make(map[string]int)
	fmt.Println("Match schedule")
	for i := 0; i < len(matches); i++ {
		fmt.Println(matches[i])
	}
	for i := 0; i < len(results); i++ {
		if results[i] == 0 {
			//away team wins
			wins[matches[i][1]]++
		} else {
			wins[matches[i][0]]++
		}
	}
	maxValue := 0
	for key, value := range wins {
		fmt.Println(key, value)
		if maxValue < value {

		}
	}

	// for i := 0; i < len(matches); i++ {
	// 	//check and insert in the array
	// 	if search(teams, matches[i][0]) == false {
	// 		teams = append(teams, matches[i][0])
	// 	}
	// 	if search(teams, matches[i][1]) == false {
	// 		teams = append(teams, matches[i][1])
	// 	}
	// }

	// for i := 0; i < len(teams); i++ {
	// 	fmt.Println(teams[i])
	// }
	// wins := make([]int, len(teams))
	// for i := 0; i < len(matches); i++ {
	// 	index := -1
	// 	if results[i] == 0 {
	// 		index = returnIndex(teams, matches[i][1])

	// 	} else {
	// 		index = returnIndex(teams, matches[i][0])
	// 	}
	// 	wins[index]++
	// }
	// fmt.Println(teams[findLargest(wins)])
}

func search(arr []string, target string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func returnIndex(arr []string, target string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
func findLargest(arr []int) int {
	maxPos := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			maxPos = i + 1
		}
	}
	return arr[maxPos]
}
