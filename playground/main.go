package main

import (
	"fmt"
	"sort"
)

func solution(nums1 []int, nums2 []int) float64 {
    mergedArr := append(nums1, nums2...)
    sort.Slice(mergedArr, func(i, j int) bool {
        return mergedArr[i] < mergedArr[j]
    })
    if len(mergedArr)%2 == 0 {
        //짝수개
        return float64(mergedArr[len(mergedArr)/2-1] + mergedArr[len(mergedArr)/2]) / 2
    } else {
        return float64(mergedArr[len(mergedArr)/2])
    }
}

func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([]int{1,2}, []int{3,4}))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
