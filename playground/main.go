package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func solution(nums []int, target int) int {
    sort.Slice(nums, func(i int, j int) bool {
        return nums[i] < nums[j]
    })
    minDiff := math.MaxInt64
    diffMap := make(map[int]int) // diff - 3sum
    for i:=0; i < len(nums)-2; i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        start, end, twoSumTarget :=  i+1, len(nums)-1, target-nums[i]
        for start < end {
            twoSum := nums[start] + nums[end]
            if twoSum == twoSumTarget {
                return target
            } else if twoSum < twoSumTarget {
                diff := abs(twoSumTarget-twoSum)
                if minDiff > diff {
                    minDiff = diff
                    diffMap[minDiff] = twoSum + nums[i]
                }
                start++
            } else if twoSum > twoSumTarget {
                diff := abs(twoSumTarget-twoSum)
                if minDiff > diff {
                    minDiff = diff
                    diffMap[minDiff] = twoSum + nums[i]
                }
                end-- 
            }
        }
    }
    return diffMap[minDiff]
    
}


func main() {
	//var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
	fmt.Println(solution([]int{-1000,-1000,-1000}, 10000))
	//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
	//     [[0, 4], [0, 1], [2, 3]] => 2

}
