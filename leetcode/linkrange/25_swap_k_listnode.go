//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。 
//
// k 是一个正整数，它的值小于或等于链表的长度。 
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 
//
// 示例： 
//
// 给你这个链表：1->2->3->4->5 
//
// 当 k = 2 时，应当返回: 2->1->4->3->5 
//
// 当 k = 3 时，应当返回: 3->2->1->4->5 
//
// 
//
// 说明： 
//
// 
// 你的算法只能使用常数的额外空间。 
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
// 
// Related Topics 链表 
// 👍 718 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.

 */
package linkrange

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	head, _ = swapNode(head, k)
	return head
}
func swapNode(node *ListNode, k int) (*ListNode, *ListNode) {
	if node == nil {
		return nil, nil
	}
	head := node
	i := 1
	for node.Next != nil && i < k {
		next := node.Next
		node.Next = next.Next
		next.Next = head
		head = next
		i++
	}
	if i != k {
		return swapNode(head, i)
	}
	if node != nil {
		node.Next, node = swapNode(node.Next, k)
	}
	return head, node
}

//leetcode submit region end(Prohibit modification and deletion)
