package main

import (
	"fmt"
)

func solution(s string) int {
	start, end, maxLen := 0, 0, 1
	characters := make(map[rune]bool)
	if len(s) == 0 {
		return 0
	}
	characters[rune(s[0])] = true
	for end < len(s)-1 {
		end++
		endChar := s[end]
		if characters[rune(endChar)] {
			for {
				if s[start] == endChar {
					start++
					break
				}
				characters[rune(s[start])] = false
				start++
			}
		} else {
			characters[rune(endChar)] = true
		}
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
		fmt.Println(start, end, maxLen)
	}
	return maxLen
}

func main() {
	//var sequence [][]int = [][]int{{0,4}, {1,2}, {1,3}, {3,4}}
	fmt.Println(solution("tmmzuxt"))
	//     [[0, 4], [1, 2], [1, 3], [3, 4]] => 2
	//     [[0, 4], [0, 1], [2, 3]] => 2

}
