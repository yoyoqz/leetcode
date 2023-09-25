package arrray

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxA := 0

	for l < r {
		maxA = max(maxA, (r-l)*min(height[r], height[l]))

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return maxA
}
