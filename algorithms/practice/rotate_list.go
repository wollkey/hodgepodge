package practice

type ListNode struct {
	Value int
	Next  *ListNode
}

var Lists = ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
var Lists2 = ListNode{1, &ListNode{2, nil}}
var Lists3 = ListNode{1, nil}

// RotateRight
// O(n)
func RotateRight(head *ListNode, k int) *ListNode {
	var stack []*ListNode
	start := head

	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	if len(stack) <= 1 {
		return start
	}

	for k >= len(stack) {
		k = k - len(stack)
	}

	if k == 0 {
		return start
	}

	end := stack[len(stack)-k-1]
	end.Next = nil

	tail := stack[len(stack)-1]
	tail.Next = start

	return stack[len(stack)-k]
}

func rotateRightOptimized(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Подсчитываем длину списка и находим последний узел
	oldTail := head
	n := 1
	for oldTail.Next != nil {
		oldTail = oldTail.Next
		n++
	}

	// Соединяем хвост с головой, формируя кольцо
	oldTail.Next = head

	// Находим новый хвост: (n - k % n - 1) узел и новую голову: (n - k % n) узел
	newTail := head
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	// Разрываем кольцо, устанавливая новый хвост
	newTail.Next = nil

	return newHead
}
