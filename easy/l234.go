package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev := &ListNode{-1, nil}

	for slow != nil {
		slowNext := slow.Next
		slow.Next = prev
		prev = slow
		slow = slowNext
	}

	first := head
	second := prev

	for second.Val != -1 {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}
