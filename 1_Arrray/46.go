package arrray

func permute(nums []int) [][]int {
	var arr []int
	var res [][]int
	used := make([]bool, len(nums))

	dfs(nums, arr, &res, used)

	return res
}

func dfs(nums []int, arr []int, res *[][]int, used []bool) {
	if len(nums) == len(arr) {
		tmp := append([]int{}, arr...)
		*res = append(*res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			arr = append(arr, nums[i])
			dfs(nums, arr, res, used)
			arr = arr[:len(arr)-1]
			used[i] = false
		}
	}
}
