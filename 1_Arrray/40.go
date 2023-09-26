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
		//小剪支
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		//大剪支
		if candidates[i] > target {
			break
		}

		//arr = append(arr, candidates[i])
		// i + 1		//不是start+1
		//append 可以不用再做删除 arr 操作
		dfs(candidates, i+1, target-candidates[i], append(arr, candidates[i]), res)
		//arr = arr[:len(arr)-1]
	}

}

//递归
/*

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(nums []int, target int) [][]int {
	ret := [][]int{}
	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		} else if target < n {
			break
		} else if target == n {
			ret = append(ret, []int{n})
			continue
		}
		for _, v := range dfs(nums[i+1:], target-n) {
			ret = append(ret, append([]int{n}, v...))
		}
	}
	return ret
}

作者：ukcuf
链接：https://leetcode.cn/problems/combination-sum-ii/solutions/315570/shuang-100de-jian-ji-dfsdi-gui-jie-by-ukcuf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


*/
