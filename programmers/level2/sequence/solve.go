package sequence

import "math"


func solution(sequence []int, k int) []int {
    seqLen := int(math.Pow(10,6))
    startIdx, endIdx :=int(math.Pow(10,6)) , 0
    for i:= len(sequence)-1; i >= 0; i-- {
        sum := sequence[i]
        if sum == k {
            if i < startIdx {
                startIdx, endIdx = i, i
                seqLen = 0
            }
        }
        for j:= i-1; j>= 0; j-- {
            sum += sequence[j]
            if i-j > seqLen {
                break
            } else if sum > k { 
                break
            } else if sum == k { 
                if i - j < seqLen {
                    startIdx = j
                    endIdx = i
                } else if i - j == seqLen {
                    if startIdx > j {
                        startIdx = j
                        endIdx = i
                    }
                }
            }

        }
    }
    return []int{startIdx, endIdx}
}