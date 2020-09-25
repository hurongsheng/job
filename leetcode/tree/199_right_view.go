//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。 
//
// 示例: 
//
// 输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
// 
// Related Topics 树 深度优先搜索 广度优先搜索 
// 👍 323 👎 0

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

func rightSideView(root *TreeNode) []int {
	return rightSideViewLevel(root, 0, make([]int, 0))
}

func rightSideViewLevel(root *TreeNode, level int, res [] int) []int {
	if root == nil {
		return res
	}
	if level >= len(res) {
		res = append(res, root.Val)
	} else {
		res[level] = root.Val
	}
	return rightSideViewLevel(root.Right, level+1, rightSideViewLevel(root.Left, level+1, res))

}

//leetcode submit region end(Prohibit modification and deletion)
