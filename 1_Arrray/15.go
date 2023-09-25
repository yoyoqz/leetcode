package arrray

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		//特殊情况
		if nums[i] > 0 {
			break
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--

			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}

		}
	}

	return res
}
