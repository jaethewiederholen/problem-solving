package searchInsert

func BasicBinarySearch(nums []int, target int) int {
    left, mid, right := 0, 0, len(nums)-1
    for left <= right {
        mid = (left + right) /2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid+1
        } else if nums[mid] > target {
            right = mid-1
        }
    }
    return right+1

}