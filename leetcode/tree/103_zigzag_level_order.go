//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。 
//
// 例如： 
//给定二叉树 [3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回锯齿形层次遍历如下： 
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
// 
// Related Topics 栈 树 广度优先搜索 
// 👍 264 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	right := true
	stack := []*TreeNode{root}
	res := make([][]int, 0)
	for len(stack) > 0 {
		tmp := make([]int, 0)
		stackTmp := make([]*TreeNode, 0)
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] == nil {
				continue
			}
			tmp = append(tmp, stack[i].Val)
			if right {
				stackTmp = append(stackTmp, stack[i].Left, stack[i].Right)
			} else {
				stackTmp = append(stackTmp, stack[i].Right, stack[i].Left)
			}
		}
		right = !right
		stack = stackTmp
		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
