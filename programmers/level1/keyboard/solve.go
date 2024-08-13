package Keyboard

import "strings"

func Solution(keymap []string, targets []string) []int {
    //알파벳 별 최소 자판수
    minPerLetter := make(map[string]int)
    
    for _, key := range keymap {
        letters := strings.Split(key, "")
        for i, l := range letters {
            if minPerLetter[l] == 0 {
                minPerLetter[l] = i+1
            } else if minPerLetter[l] > i+1 {
                minPerLetter[l] = i+1
            }
        }
    }
    result := make([]int, len(targets))
    
    for i, target := range targets {
        res := 0
        letters := strings.Split(target, "")
        for _, t := range letters {
            if minPerLetter[t] == 0 {
                res = -1
                break
            } else {
                res += minPerLetter[t]
            }
        }
        result[i] = res
    }
    return result
}