package arrray

/*
双指针

*/

func removeElement(nums []int, val int) int {
	l := 0

	//注意为len(nums)
	r := len(nums)

	for l < r {
		if nums[l] == val {
			nums[l] = nums[r-1]
			r--
		} else {
			l++
		}
	}

	return l
}
