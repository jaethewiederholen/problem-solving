package sequence

import "math"

func solution_fullSearch(sequence []int, k int) []int {
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
