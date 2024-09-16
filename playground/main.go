package main

import (
	"fmt"
)

func solution(nums []int, target int) int {
    left, mid, right := 0, 0, len(nums)-1
    for left <= right {
        mid = (left + right) /2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid+1
        } else if nums[mid] > target {
            right = mid-1
        }
    }
    return right+1
}


func main() {
	//var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
	fmt.Println(solution([]int{1,3,5,6}, 2))
	//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
	//     [[0, 4], [0, 1], [2, 3]] => 2

}
