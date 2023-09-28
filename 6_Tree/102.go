package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个队列
// 循环队列取值
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode

	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		var arr []int

		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			arr = append(arr, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, arr)

	}

	return res
}
