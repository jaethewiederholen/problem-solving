package sequence

import "math"

func Solution_fullSearch(sequence []int, k int) []int {
	seqLen := int(math.Pow(10, 6))
	startIdx, endIdx := int(math.Pow(10, 6)), 0
	for i := 0; i <= len(sequence)-1; i++ {
		sum := sequence[i]
		if sum == k {
			startIdx, endIdx = i, i
			break
		}
		for j := i + 1; j <= len(sequence)-1; j++ {
			sum += sequence[j]
			if j-i > seqLen {
				break
			}
			if sum == k {
				if j-i < seqLen {
					startIdx, endIdx = i, j
					seqLen = j - i
					break
				}
			}
		}
	}
	return []int{startIdx, endIdx}
}

func Solution_twoPointer(sequence []int, k int) []int {
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