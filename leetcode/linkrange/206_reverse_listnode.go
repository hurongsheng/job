//反转一个单链表。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表 
// 👍 1216 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package linkrange

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	parent := &ListNode{0, head}
	for head.Next != nil {
		tmp := head.Next
		head.Next = head.Next.Next
		tmp.Next = parent.Next
		parent.Next = tmp
	}
	return parent.Next

}

//leetcode submit region end(Prohibit modification and deletion)
