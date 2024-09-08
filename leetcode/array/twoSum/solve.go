package twoSum

func Solution(nums []int, target int) []int {
    type aInfo struct {
        idx int
    }
    bDic := make(map[int]*aInfo)
    
    ans := make([]int, 2)
   
    for idx, num := range nums {
        if bDic[num] != nil {
            ans[0] = bDic[num].idx
            ans[1] = idx
            break
        } else {
            bDic[target-num] = &aInfo{idx: idx}
        }

    }
    return ans
}