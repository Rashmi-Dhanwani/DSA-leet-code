/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy

    // Advance fast by n+1 steps, so slow lands on the node *before* the target.
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Move both until fast falls off the end.
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next // unlink the nth-from-end node
    return dummy.Next
}