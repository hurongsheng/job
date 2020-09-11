//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。 
//
// 示例 1: 
//
// 输入: 1->2->3->3->4->4->5
//输出: 1->2->5
// 
//
// 示例 2: 
//
// 输入: 1->1->1->2->3
//输出: 2->3 
// Related Topics 链表 
// 👍 361 👎 0

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

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	flag := false
	for head.Next != nil {
		if head.Next.Val == head.Val {
			head.Next = head.Next.Next
			if head.Next == nil {
				return nil
			}
			flag = true
		} else {
			if flag {
				return DeleteDuplicates(head.Next)
			}
			head.Next = DeleteDuplicates(head.Next)
			break
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
