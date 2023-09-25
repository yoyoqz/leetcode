package arrray

import "sort"

/*
不包含重复的数字

*/

func combinationSum2(candidates []int, target int) [][]int {
	var arr []int
	var res [][]int

	sort.Ints(candidates)

	dfs(candidates, 0, target, arr, &res)

	return res
}

func dfs(candidates []int, start int, target int, arr []int, res *[][]int) {
	if target == 0 {
		tmp := append([]int{}, arr...)
		*res = append(*res, tmp)

		//忘记加return
		return
	}

	for i := start; i < len(candidates); i++ {
		//大剪支
		if candidates[i] > target {
			break
		}

		//小剪支
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		arr = append(arr, candidates[i])
		// + 1
		dfs(candidates, start+1, target-candidates[i], arr, res)
		arr = arr[:len(arr)-1]
	}

}
