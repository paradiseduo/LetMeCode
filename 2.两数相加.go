/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers2(l1 *ListNode, l2 *ListNode, add int) *ListNode {
	if l1 == nil && l2 == nil {
		if add != 0 {
			return &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	} else if l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + add
		return &ListNode{
			Val:  sum % 10,
			Next: addTwoNumbers2(l1.Next, l2.Next, sum/10),
		}
	} else if l1 == nil {
		sum := l2.Val + add
		return &ListNode{
			Val:  (l2.Val + add) % 10,
			Next: addTwoNumbers2(nil, l2.Next, sum/10),
		}
	} else if l2 == nil {
		sum := l1.Val + add
		return &ListNode{
			Val:  (l1.Val + add) % 10,
			Next: addTwoNumbers2(l1.Next, nil, sum/10),
		}
	}
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers2(l1, l2, 0)
}
