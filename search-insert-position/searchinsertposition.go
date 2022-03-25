package searchinsertposition

func SearchInsert(nums []int, target int) int {
	left, right, pos := 0, len(nums)-1, 0

	for left <= right {
		pos = (left + right) / 2
		if nums[pos] == target {
			return pos
		}
		if target < nums[pos] {
			right = pos - 1
		} else {
			left = pos + 1
		}

	}
	return left
}
