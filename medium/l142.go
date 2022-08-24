package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	addresses := make(map[*ListNode]bool)
	ptr := head

	for ptr != nil {
		if addresses[ptr] {
			return ptr
		} else {
			addresses[ptr] = true
			ptr = ptr.Next
		}
	}
	return nil
}
