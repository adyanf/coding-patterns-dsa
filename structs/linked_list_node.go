package structs

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	node := new(LinkedListNode)
	node.Data = data
	node.Next = next
	return node
}

func InitLinkedListNode(data int) *LinkedListNode {
	node := new(LinkedListNode)
	node.Data = data
	node.Next = nil
	return node
}

func ReverseLinkedList(head *LinkedListNode) *LinkedListNode {
	var prev, next, curr *LinkedListNode = nil, nil, head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
