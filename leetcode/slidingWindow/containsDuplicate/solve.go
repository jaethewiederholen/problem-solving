package containsDuplicate

func ContainsNearbyDuplicate(nums []int, k int) bool {
	i, j := 0, 0
	contains := make(map[int]bool)
	for j < len(nums) {
		if contains[nums[j]] && j-i <= k {
			return true
		}
		if j-i > k {
			for j-i > k {
				delete(contains, nums[i])
				i++
			}
		} else {
			contains[nums[j]] = true
			j++
		}
	}
	return false
}
