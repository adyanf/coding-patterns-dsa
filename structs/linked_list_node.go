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
