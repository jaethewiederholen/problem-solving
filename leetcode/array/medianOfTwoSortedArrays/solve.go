package medianOfTwoSortedArrays

import "sort"

func BurtForceSolution(nums1 []int, nums2 []int) float64 {
    mergedArr := append(nums1, nums2...)
    sort.Slice(mergedArr, func(i, j int) bool {
        return mergedArr[i] < mergedArr[j]
    })
    if len(mergedArr)%2 == 0 {
        //짝수개
        return float64(mergedArr[len(mergedArr)/2-1] + mergedArr[len(mergedArr)/2]) / 2
    } else {
        return float64(mergedArr[len(mergedArr)/2])
    }
}

func BinarySearchSolution(nums1 []int, nums2 []int) float64 {
	//TODO
	return 0.0
}