package threeSum

import (
	"sort"
)

func twoSum(nums []int, target int) [][2]int {
	bMap := make(map[int]bool)
	var sets [][2]int
	for _, b := range nums {
		if bMap[b] {
			set := [2]int{target - b, b}
			sets = append(sets, set)
		} else {
			bMap[target-b] = true
		}
	}
	return sets
}

func solution(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	threeSumAns := make(map[[3]int]bool)

	for i := 0; i < len(nums); i++ {
		elem1 := nums[i]
		if elem1 > 0 {
			break
		}

		target := 0 - elem1
		twoSumAns := make(map[[2]int]bool)
		twoSets := twoSum(nums[i+1:], target)

		for _, set := range twoSets {
			twoSumAns[set] = true
		}

		for twoSum, _ := range twoSumAns {
			threeSum := [3]int{elem1, twoSum[0], twoSum[1]}
			threeSumAns[threeSum] = true
		}
	}
	var ans [][]int
	for threeSum, _ := range threeSumAns {
		ans = append(ans, []int{threeSum[0], threeSum[1], threeSum[2]})
	}
	return ans
}
