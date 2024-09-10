package maxArea

func TwoPointerSolution(height []int) int {
    start := 0
    end := len(height)-1
    max := 0
    for {
        if start >= end {
            break
        }

        startH := height[start]
        endH := height[end]
        x := end - start
        y := 0
        if startH > endH {
            y = endH
            end--
        } else {
            y = startH
            start++
        }
        if max < x * y {
            max = x*y
        }
    }
    return max
}