package practice

// WiggleSort
// O(n)
func WiggleSort(nums []int) []int {
	length := len(nums)

	for i := 1; i < length; i++ {
		if (i%2 == 1 && nums[i] < nums[i-1]) || (i%2 == 0 && nums[i] > nums[i-1]) {
			temp := nums[i-1]
			nums[i-1] = nums[i]
			nums[i] = temp
		}
	}

	return nums
}
