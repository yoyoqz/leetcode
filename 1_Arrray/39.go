package arrray

import "sort"

//组合

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var arr []int

	sort.Ints(candidates)

	dfs(candidates, 0, len(candidates), target, arr, &res)

	return res
}

func dfs(candidates []int, start int, length int, target int, arr []int, res *[][]int) {
	if target == 0 {
		tmp := append([]int{}, arr...)
		*res = append(*res, tmp)

		return
	}

	for i := start; i < length; i++ {
		//剪枝
		if target-candidates[i] < 0 {
			break
		}

		arr = append(arr, candidates[i])

		dfs(candidates, i, length, target-candidates[i], arr, res)

		arr = arr[:len(arr)-1]
	}
}
