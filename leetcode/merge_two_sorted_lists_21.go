package main

//You are given the heads of two sorted linked lists list1 and list2.
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//Return the head of the merged linked list.

// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.
func main() {
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}
