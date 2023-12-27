// Package doublylinkedlist implements a generic version of a doubly-linked list.
package doublylinkedlist

// DoublyLinkedList that uses a sentinel node and is circular
// TODO: add in some json marshalling and unmarshalling support one day
type DoublyLinkedList[T comparable] struct {
	head Node[T] // sentinel node
	// commenting for experiment.
	// tail Node[T] // another sentinel node, but may not need if we go circular
	len int // current list length which excludes any sentinel node(s)
}

// Node is an item in the list.
type Node[T comparable] struct {
	next, prev *Node[T]

	// reference to the list that owns this node
	list *DoublyLinkedList[T]

	// The value stored with this node.
	Value T
}

// New returns a new, initialized, empty DoublyLinkedList[T].
func New[T comparable]() *DoublyLinkedList[T] {
	return new(DoublyLinkedList[T]).Initialize()
}

func (l *DoublyLinkedList[T]) Initialize() *DoublyLinkedList[T] {
	l.head.next = &l.head
	l.head.prev = &l.head
	// I thought I'd need both head and tail nodes like textbook implementations,
	// but maybe one sentinel is enough. Commenting for experimentation
	// l.tail.prev = &l.head
	// l.tail.next = &l.head
	l.len = 0 // never count the sentinel node(s) in length.
	return l
}

// ensureInitialized will properly initialize the list if it has not been already
// this means not having to return uninitialized errors for eager callers of Adds.
func (l *DoublyLinkedList[T]) ensureInitialized() {
	if l.head.next == nil {
		l.Initialize()
	}
}

// Len returns the length of the list. Sentinel node not included in count.
func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}

// First returns the first node of the list or nil if the list is empty .
func (l *DoublyLinkedList[T]) First() *Node[T] {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Last returns the last node of the list or nil if the list is empty .
func (l *DoublyLinkedList[T]) Last() *Node[T] {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// Next returns the next node of the list or nil if the list is empty.
func (n *Node[T]) Next() *Node[T] {
	// trying without tail sentinel, so you know you are at the end
	// if head would come after the last node since this is circular
	if nextNode := n.next; n.list != nil && nextNode != &n.list.head {
		return nextNode
	}
	return nil
}

// Next returns the previous node of the list or nil if the list is empty.
func (n *Node[T]) Prev() *Node[T] {
	if prevNode := n.prev; n.list != nil && prevNode != &n.list.head {
		return prevNode
	}
	return nil
}

// insert inserts the value as a new *Node after the nodeToInsertAfter node
// length of the list is incremented and the new node created from the value is returned.
func (l *DoublyLinkedList[T]) insert(v T, nodeToInsertAfter *Node[T]) *Node[T] {
	nodeToInsert := &Node[T]{
		prev:  nodeToInsertAfter,
		next:  nodeToInsertAfter.next,
		Value: v,
		list:  l,
	}

	nodeToInsert.next.prev = nodeToInsert
	nodeToInsertAfter.next = nodeToInsert

	l.len++
	return nodeToInsert
}

// AddFirst adds a new node with the value at the beginning of the list.
func (l *DoublyLinkedList[T]) AddFirst(v T) *Node[T] {
	l.ensureInitialized()
	return l.insert(v, &l.head)
}

// AddLast adds a new node with the value at the end of the list
// (essentially before the head node where the prev always points to the last node).
func (l *DoublyLinkedList[T]) AddLast(v T) *Node[T] {
	l.ensureInitialized()
	return l.insert(v, l.head.prev)
}

// Contains performs a linear search for the given value in the list
// true if found, false if not.
func (l *DoublyLinkedList[T]) Contains(v T) bool {
	for n := l.First(); n != nil; n = n.Next() {
		if n.Value == v {
			return true
		}
	}
	return false
}
