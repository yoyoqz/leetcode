package __Arrray

// 1、declare a map to store the value as the key, store the key as the value
// 2、
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for k, v := range nums {
		if v2, isOK := m[target-v]; isOK {
			return []int{v2, k}
		}
		m[v] = k
	}

	return []int{}
}
