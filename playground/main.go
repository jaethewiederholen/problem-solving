package main

import (
	"fmt"
	"sort"
)
func solution(targets [][]int) int {
    
    sort.Slice(targets, func(i,j int) bool {
        if targets[i][0] < targets[j][0] {
            return true
        } else if targets[i][0] == targets[j][0] {
            return targets[i][1] > targets[j][1]
        } else {
            return false
        }
    })
    ans := 1
    start := targets[0][0]
    end := targets[0][1]
    for _, target := range targets {
        if end > target[0] {
            // 겹치면
            if target[0] > start {
                start = target[0]
            }
            if target[1] < end{
                end = target[1]
            }
        } else {
            ans+=1
            start = target[0]
            end= target[1]
        }
    }
    return ans
}


func main() {
    var targets [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution(targets))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
