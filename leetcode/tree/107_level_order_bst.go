//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历） 
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
// 返回其自底向上的层次遍历为： 
//
// [
//  [15,7],
//  [9,20],
//  [3]
//]
// 
// Related Topics 树 广度优先搜索 
// 👍 332 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	return levelOrderBottoms([]*TreeNode{root})
}

func levelOrderBottoms(nodes []*TreeNode) [][]int {
	if len(nodes) == 0 {
		return make([][]int, 0)
	}
	next := make([]*TreeNode, 0)
	vals := make([]int, 0, len(nodes))
	for _, node := range nodes {
		vals = append(vals, node.Val)
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}
	return append(levelOrderBottoms(next), vals)
}

func levelOrderBottom2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		level := make([]int, 0)
		next := make([]*TreeNode, 0, 2*len(nodes))
		for _, node := range nodes {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		nodes = next
		res = append([][]int{level}, res...)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
