//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。 
//
// 说明: 
//1 ≤ m ≤ n ≤ 链表长度。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL 
// Related Topics 链表 
// 👍 508 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package linkrange

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	i := 0
	head = &ListNode{0, head}
	ret := head
	var mParent *ListNode
	for head.Next != nil {
		i++
		if i < m {
			head = head.Next
			continue
		}
		if i == m {
			mParent = head
			head = head.Next
			continue
		}
		if i <= n {
			move := head.Next
			head.Next = head.Next.Next
			move.Next = mParent.Next
			mParent.Next = move
		} else {
			break
		}
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)
