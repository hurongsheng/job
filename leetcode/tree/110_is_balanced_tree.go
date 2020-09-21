//给定一个二叉树，判断它是否是高度平衡的二叉树。 
//
// 本题中，一棵高度平衡二叉树定义为： 
//
// 
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。 
// 
//
// 示例 1: 
//
// 给定二叉树 [3,9,20,null,null,15,7] 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回 true 。 
// 
//示例 2: 
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4] 
//
//        1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// 
//
// 返回 false 。 
//
// 
// Related Topics 树 深度优先搜索 
// 👍 480 👎 0

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

func isBalanced(root *TreeNode) bool {
	ok, _ := treeHightEq(root)
	return ok
}
func treeHightEq(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	ok, l := treeHightEq(node.Left)
	if !ok {
		return ok, l
	}
	ok, r := treeHightEq(node.Right)
	if !ok {
		return ok, r
	}
	if l >= r {
		return l-r <= 1, l + 1
	} else {
		return r-l <= 1, r + 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
