package main

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
//Constraints:
//
//The number of nodes in the list is in the range [0, 300].
//-100 <= Node.val <= 100
//The list is guaranteed to be sorted in ascending order.

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	if head == nil {
		return nil
	}

	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
