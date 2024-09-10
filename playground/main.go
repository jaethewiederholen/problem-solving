package main

import (
	"fmt"
	"sort"
)

func solution(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    fmt.Println(nums)
    var ans [][]int
    for i := 0; i < len(nums); i ++ {
        // if nums[i] == 0 {
        //     if len(nums) -i > 2 && nums[i+1] == 0 && nums[i+2] == 0 {
        //         ans = append(ans, []int{0,0,0})
        //     }
        // }
    SkipDup1 : 
        if i > 0 && nums[i-1] == nums[i] {
            if i == len(nums)-1 { return ans}
            i++
            goto SkipDup1
        }
        target := nums[i]
        if target > 0 {break}
        elems := make([]int , 3)
        elems[0] = target
        
        sum := 0-target
        sumDic := make(map[int]bool) // elem2
        fmt.Println("target---------", target)
        for j := i+1; j< len(nums); j++ {
            elem1 := nums[j]
            elem2 := sum-elem1
            if sumDic[elem1] {
                elems[2] = elem2
                elems[1] = elem1
                fmt.Println(i, j, elem1, elem2)
                ans = append(ans, []int{target, elem1, elem2})
            } else {
                sumDic[elem2] = true
               
                fmt.Println(sumDic)
                SkipDup2 : 
            if j < len(nums)-2 && nums[j] == nums[j+2] {
                j++
                goto SkipDup2
            }
            }
        }
    }
    return ans
}

func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([]int{-2,0,0,2,2}))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
