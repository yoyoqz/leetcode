package arrray

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return sum
			}
		}
	}

	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -1 * a
}
