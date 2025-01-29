package linked_list_in_place_manipulation

import "github.com/adyanf/coding-patterns-dsa/structs"

// The in-place manipulation of a linked list pattern allows us to modify a linked list without using any additional memory.
// In-place refers to an algorithm that processes or modifies a data structure using only the existing memory space,
// without requiring additional memory proportional to the input size.
// This pattern is best suited for problems where we need to modify the structure of the linked list, i.e., the order in which nodes are linked together.
// How can we implement the in-place reversal of nodes so that no extra space is used?
// We iterate over the linked list while keeping track of three nodes: the current node, the next node, and the previous node.
// Keeping track of these three nodes enables us to efficiently reverse the links between every pair of nodes.
// This in-place reversal of a linked list works in O(n) time and consumes only O(1) space.
// Use this pattern when these conditions are fulfilled:
// - Linked list restructuring: The input data is given as a linked list, and the task is to modify its structure without modifying the data of the individual nodes.
// - In-place modification: The modifications to the linked list must be made in place, that is, we’re not allowed to use more than O(1) additional space.

// SwapPairs swaps every two adjacent nodes of the linked list. After the swap, return the head of the linked list.
func SwapPairs(head *structs.LinkedListNode) *structs.LinkedListNode {
	// if the linked list length is less than 2 then return the head
	if head == nil || head.Next == nil {
		return head
	}

	// because min length of the linked list is 2, we can assume the new head will be the next of the head
	newHead := head.Next

	// keep track of current node and previous node
	prev := (*structs.LinkedListNode)(nil)
	current := head
	for current != nil && current.Next != nil {
		// 2 -> 1(prev) -> 3(current) -> 4 -> 5 to 2 -> 1 -> 4 -> 3(prev) -> 5(current)
		// connect 3 -> 5
		// connect 4 -> 3
		// connect 1 -> 4
		next := current.Next
		current.Next = next.Next
		next.Next = current
		if prev != nil {
			prev.Next = next
		}
		// move prev to current node
		prev = current
		// move current to current next node
		current = current.Next
	}

	return newHead
}

// ReverseBetween reverse the nodes of the list from left to right, given a singly linked list with n nodes and left and right positions.
func ReverseBetween(head *structs.LinkedListNode, left int, right int) *structs.LinkedListNode {
	// if head node is empty or if left equal to right (then no node needed to be reversed), return head directly
	if head == nil || left == right {
		return head
	}

	newHead := head

	// move the pointer until reach the left
	prevLeft := (*structs.LinkedListNode)(nil)
	nodeLeft := head
	i := 1
	for i < left {
		prevLeft = nodeLeft
		nodeLeft = nodeLeft.Next
		i++
	}

	// adjust the pointer next for each node which index between left and right
	prev := (*structs.LinkedListNode)(nil)
	current := nodeLeft
	next := nodeLeft.Next
	for i <= right {
		current.Next = prev

		prev = current
		current = next
		if current != nil {
			next = current.Next
		}
		i++
	}

	// connect the prev left node to right node
	if prevLeft != nil {
		prevLeft.Next = prev
	}
	// connect left node to next right node (which is pointed by current)
	nodeLeft.Next = current

	// if the left equal to 1 then the head will change to the right node, which is the prev node
	if left == 1 {
		newHead = prev
	}
	return newHead
}

// ReverseBetweenV2 reverse the nodes of the list from left to right, given a singly linked list with n nodes and left and right positions.
// this method using dummy node as helper and in each iteration we reverse all the connection completely
func ReverseBetweenV2(head *structs.LinkedListNode, left int, right int) *structs.LinkedListNode {
	// if head node is empty or if left equal to right (then no node needed to be reversed), return head directly
	if head == nil || left == right {
		return head
	}

	// introduce dummy node to help with edge cases when we need to reverse from the head
	dummy := structs.NewLinkedListNode(0, nil)
	dummy.Next = head
	prev := dummy

	// move the prev node to node before left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// curr node will start at left node
	curr := prev.Next

	// iterate for right - left times, in each iteration
	// set the next node from curr next
	// point the curr next to next node next
	// point the next node next to prev node next
	// point the prev node next to next node
	for i := 0; i < right-left; i++ {
		nextNode := curr.Next
		curr.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}

// ReorderList reorder the list as if it were folded on itself
// L0 -> L1 -> L2 -> Ln-2 -> Ln-1 -> Ln ==> L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2
func ReorderList(head *structs.LinkedListNode) {
	// if length of the linked list is less than 3 then return head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// search the middle part of the linked list
	// when the fast pointer reach the end of linked list
	// the slow pointer will be at the start of second half of the linked list
	// we need to keep the prev slow too, since it will be the end of the first half of linked list
	slow, fast, prevSlow := head, head, (*structs.LinkedListNode)(nil)
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the linked list
	var prev, next, curr *structs.LinkedListNode = nil, nil, slow
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	prevSlow.Next = prev

	// Reorder list by moving the node in fast pointer between the node in slow pointer
	prevFast := prevSlow
	fast = prevSlow.Next
	slow = head
	for slow != prevFast {
		nextSlow := slow.Next
		nextFast := fast.Next
		slow.Next = fast
		fast.Next = nextSlow
		prevFast.Next = nextFast
		slow = nextSlow
		fast = nextFast
	}
}
