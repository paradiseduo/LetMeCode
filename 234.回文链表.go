/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	n := head
	var arr []int
	for {
		if n == nil {
			break
		}
		arr = append(arr, n.Val)
		n = n.Next
	}
	l := len(arr)
	m := l%2
	for i := 0; i < l; i++ {
		if i == ((l-m)/2) {
			return true
		}
		if arr[i] != arr[l-(i+1)] {
			return false
		}
	}
	return true
}
