package arrray

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		//去重1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//太大直接返回
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		for j := i + 1; j < len(nums)-2; j++ {
			//去重2
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			k := j + 1
			m := len(nums) - 1
			for k < m {
				sum := nums[i] + nums[j] + nums[k] + nums[m]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
					////去重3
					for k < m && nums[k] == nums[k+1] {
						k++
					}
					for k < m && nums[m] == nums[m-1] {
						m--
					}

					m--
					k++
				} else if sum > target {
					m--
				} else if sum < target {
					k++
				}
			}
		}
	}

	return res
}
