package longestSubstring

func LengthOfLongestSubstring(s string) int {
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
	}
	return maxLen
}
