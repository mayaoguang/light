package generic

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func (slf *ListNode[T]) ReverseList(head *ListNode[T]) (r *ListNode[T]) {
	if head == nil {
		return nil
	}
	r = new(ListNode[T])
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = r
		r = cur
		cur = next
	}
	return r
}
