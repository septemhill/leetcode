/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	c := head
	prev := (*ListNode)(nil)

	for c != nil {
		n := c.Next
		c.Next = prev
		prev = c
		c = n
	}

	return prev
}
