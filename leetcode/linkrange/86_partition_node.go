//给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。 
//
// 你应当保留两个分区中每个节点的初始相对位置。 
//
// 示例: 
//
// 输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5
// 
// Related Topics 链表 双指针 
// 👍 250 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package linkrange

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var xPoint *ListNode
	head = &ListNode{0, head}
	ret := head
	for head.Next != nil {
		if head.Next.Val >= x {
			if xPoint == nil {
				xPoint = head
			}
			head = head.Next
		} else {
			if xPoint == nil {
				head = head.Next
			} else {
				tmp := head.Next
				head.Next = head.Next.Next
				tmp.Next = xPoint.Next
				xPoint.Next = tmp
				xPoint = xPoint.Next
			}
		}
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)
