/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	first := list1
	for i := 0; i < a-1; i++ {
		first = first.Next
	}
	second := first
	for i := 0; i < (b - a + 2); i++ {
		second = second.Next
	}
	first.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = second
	return list1
}
