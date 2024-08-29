package main

import (
	"fmt"
)

func solution(sequence []int, k int) []int {
    start, end, sum := len(sequence)-1, len(sequence)-1, sequence[len(sequence)-1]
    ans := make([]int,2)
    //if sum == k { return []int{start, end}}
    for {
        if start > end { break }
        if sum == k {
            ans[0] , ans[1] = start, end
            for {
                if start > 0 && end > 0 {
                    start -= 1
                    sum += sequence[start]
                    sum -= sequence[end]
                    end -= 1
                    if k == sum {
                        ans[0] , ans[1] = start, end
                    } else { return ans}
                } else {
                    return ans
                }
            }
        } else if sum < k {
            start -= 1 
            sum += sequence[start]
        } else {
            sum -= sequence[end]
            end -= 1
        }
    }
    return ans
}


func main() {
    //var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
    fmt.Println(solution([]int{7,5,5,1,1,50,50}, 100))
//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
//     [[0, 4], [0, 1], [2, 3]] => 2
	
}
