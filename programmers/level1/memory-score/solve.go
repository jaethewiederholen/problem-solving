package memoryScore

func Solution(name []string, yearning []int, photo [][]string) []int {
    scoreMap:= make(map[string]int)
    res := make([]int,len(photo))
    for i, n := range name {
        scoreMap[n] = yearning[i]
    }
    for i, p := range photo {
        for _, j := range p {
            res[i] += scoreMap[j]
        }
    }
    return res
}