package interceptSystem

import (
	"sort"
)
func Solution(targets [][]int) int {
    
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
