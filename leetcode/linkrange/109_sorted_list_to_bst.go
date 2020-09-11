//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。 
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。 
//
// 示例: 
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
// 
// Related Topics 深度优先搜索 链表 
// 👍 362 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.

 */
/**
 * Definition for a binary tree node.

 */
package linkrange

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	return BuildTree(list)
}
func BuildTree(node []*ListNode) *TreeNode {
	if len(node) == 0 {
		return nil
	}
	if len(node) == 1 {
		return &TreeNode{
			Val: node[0].Val,
		}
	}
	return &TreeNode{
		Val:   node[len(node)/2].Val,
		Left:  BuildTree(node[:len(node)/2]),
		Right: BuildTree(node[len(node)/2+1:]),
	}
}

//leetcode submit region end(Prohibit modification and deletion)
