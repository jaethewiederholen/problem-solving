package mostGivenPresent

import "strings"

func Solution(friends []string, gifts []string) int {
    ans := 0 
    present_rate := make(map[string]int)
    fromToCount := make(map[string]map[string]int)
    
    for _, gift := range gifts {
        arr:=strings.Split(gift, " ")
        present_rate[arr[0]] +=1
        present_rate[arr[1]] -=1
        if _, ok := fromToCount[arr[0]]; !ok {
            fromToCount[arr[0]] = make(map[string]int)
        }
        fromToCount[arr[0]][arr[1]] += 1
    }
    
    for _, from := range friends {
        total := 0
        for _, to := range friends {
            if from == to {
                continue
            }
            if fromToCount[from][to] > fromToCount[to][from] {
                total +=1
            } else if fromToCount[from][to] == fromToCount[to][from] {
                if present_rate[from] > present_rate[to] {
                    total +=1
                }
            }
        }
        if ans < total {
            ans = total
        }
    }

    return ans
}