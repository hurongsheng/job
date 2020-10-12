//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。 
//
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 进阶： 
//
// 你能在线性时间复杂度内解决此题吗？ 
//
// 
//
// 示例: 
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7] 
//解释: 
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10^5 
// -10^4 <= nums[i] <= 10^4 
// 1 <= k <= nums.length 
// 
// Related Topics 堆 Sliding Window 
// 👍 564 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package slicerange

import "fmt"

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 1 || k <= 1 {
		return nums
	}
	stack := []int{0}
	for i := 1; i < k; i++ {
		stack = createStack(nums, stack, k, i)
		fmt.Printf("%+v\n", stack)
	}
	res := make([]int, 0)
	res = append(res, nums[stack[0]])
	for i := k; i < len(nums); i++ {
		if stack[0] < i {
			stack = stack[1:]
		}
		stack = createStack(nums, stack, k, i)
		fmt.Printf("%+v\n", stack)
		res = append(res, nums[stack[0]])
	}
	return res
}

func createStack(nums, stack []int, k, i int) []int {
	flag := false
	for j, v := range stack {
		if nums[i] > nums[v] {
			flag = true
			stack = append(append(stack[:j], i), stack[j:]...)
			break
		}
	}
	if !flag {
		stack = append(stack, i)
	}
	if len(stack) > k {
		return stack[:k]
	}
	return stack
}

//leetcode submit region end(Prohibit modification and deletion)
