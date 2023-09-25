//Next Permutation Go Medium 34.2%
func nextPermutation(nums []int) {
	stack = make([]bool, len(nums))
	var arr []int
	var res [][]int

	dfs(nums, arr, &res, stack)

}

func dfs(nums []int, arr []int, res *[][]int, stack []bool) {
	if len(nums) == len(arr) {
		tmp := append([]int{}, arr...)
		*res = apppend(*res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if stack[i] {
			continue
		}

		stack[i] = true
		arr = append(arr, nums[i])

		dfs(nums, arr, res)

		stack[i] = false
		arr = arr[:len(nums)-1]
	}

}