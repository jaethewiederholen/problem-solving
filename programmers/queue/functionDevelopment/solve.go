package functionDevelopment

func Solution(progresses []int, speeds []int) []int {
    days := make([]int, len(progresses))
    for i, _ := range progresses {
        day := (100 - progresses[i])/speeds[i]
        if (100 - progresses[i])%speeds[i] > 0 {
            day++
        }
        days[i] = day
    }
    var ans []int
    max := 0
    for _, d := range days {
        if d > max {
            ans = append(ans, 1)
            max = d 
        } else {
            ans[len(ans)-1]++
        }
    }
    return ans
}