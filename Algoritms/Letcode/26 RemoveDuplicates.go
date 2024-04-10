package letcode

func RemoveDuplicates(nums []int) int {
	n_map := make(map[int]bool, len(nums))
	var i int
	for i = 0; i < len(nums); i++ {
		if !n_map[nums[i]] {
			n_map[nums[i]] = true
		} else {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}

	return len(nums)
}
