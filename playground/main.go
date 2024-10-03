package main

import (
	"fmt"
)

func solution(nums []int, k int) bool {
	i, j := 0, 0
	contains := make(map[int]bool)
	for j < len(nums) {
		if contains[nums[j]] && j-i <= k {
			fmt.Println(i, j)
			return true
		}
		if j-i > k {
			for j-i > k {
				delete(contains, nums[i])
				i++
			}
		} else {
			contains[nums[j]] = true
			j++
		}
		fmt.Println(i, j)
	}
	return false
}

func main() {
	//var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
	fmt.Println(solution([]int{1, 2, 3, 1, 2, 3}, 2))
	//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
	//     [[0, 4], [0, 1], [2, 3]] => 2

}
