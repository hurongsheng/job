//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。 
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。 
//
// 
//
// 进阶： 
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。 
//
// 
//
// 示例： 
//
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
// 
// Related Topics 链表 
// 👍 274 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package linkrange

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	stack := make([]*ListNode, 0)
	for l1 != nil {
		stack = append(stack, l1)
		l1 = l1.Next
	}
	stackL2 := make([]*ListNode, 0)
	for l2 != nil {
		stackL2 = append(stackL2, l2)
		l2 = l2.Next
	}
	if len(stack) > len(stackL2) {
		stackL2 = append(make([]*ListNode, len(stack)-len(stackL2)), stackL2...)
	} else if len(stack) < len(stackL2) {
		stack = append(make([]*ListNode, len(stackL2)-len(stack)), stack...)
	}
	var this *ListNode
	flag := 0
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == nil {
			this=&ListNode{stackL2[i].Val + flag, this}
		}else if stackL2[i] == nil{
			this=&ListNode{stack[i].Val + flag, this}
		}else{
			this = &ListNode{stack[i].Val + stackL2[i].Val + flag, this}
		}
		if this.Val >= 10 {
			this.Val -= 10
			flag = 1
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		return &ListNode{1, this}
	} else {
		return this
	}

}

//leetcode submit region end(Prohibit modification and deletion)
