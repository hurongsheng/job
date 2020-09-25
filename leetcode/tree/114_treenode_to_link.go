//给定一个二叉树，原地将它展开为一个单链表。 
//
// 
//
// 例如，给定二叉树 
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6 
//
// 将其展开为： 
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6 
// Related Topics 树 深度优先搜索 
// 👍 558 👎 0

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

func flatten(root *TreeNode) {
	list := preorderTraversal2(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal2(root *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal2(root.Left)...)
		list = append(list, preorderTraversal2(root.Right)...)
	}
	return list
}

//leetcode submit region end(Prohibit modification and deletion)
