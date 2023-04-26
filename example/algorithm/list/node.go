package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// SortList 链表排序
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	left := SortList(head)
	right := SortList(tail)
	return mergeTwoLists(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

// mergeTwoLists 合并两个有序数组
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
		} else {
			head.Next, l2 = l2, l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	} else if l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return dummy.Next
}
