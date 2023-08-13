package doublyLinkedList

// Node represents a node in the doubly linked list.
type Node struct {
	data int
	prev *Node
	next *Node
}

// DoublyLinkedList represents the doubly linked list.
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert inserts a new node with the given data at the front of the doubly linked list.
// (插入到链结最前面)
func (dll *DoublyLinkedList) Insert(data int) {
	// Declare a new node first.
	newNode := &Node{data: data}

	if dll.head == nil {
		// If the linked list is empty, set the new node as both the head and tail.
		// (如果链结是空就新增第一结点)
		dll.head = newNode
		dll.tail = newNode
	} else {
		// If the linked list is not empty, insert the new node to the current head and update the head pointer.
		// (新增到链结的开头)
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

// Append adds a new node to the end of the list.
// (加入到链结的最后面)
func (dll *DoublyLinkedList) Append(data int) {
	// Declare a new node first.
	newNode := &Node{data: data}

	if dll.head == nil {
		// If the linked list is empty, set the new node as both the head and tail.
		// (如果链结是空就新增第一结点)
		dll.head = newNode
		dll.tail = newNode
	} else {
		// If the linked list is not empty, append the new node to the current tail and update the tail pointer.
		// (新增到链结的尾端)
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

// Reverse reverses the order of nodes in the doubly linked list.
func (dll *DoublyLinkedList) Reverse() {
	current := dll.head
	var previous *Node

	for current != nil {
		next := current.next
		current.next = previous
		current.prev = next
		previous = current
		current = next
	}

	dll.head, dll.tail = dll.tail, dll.head
}

// ForwardList prints the elements of the list from head to tail.
// (连 头部 开始收集资料)
func (dll *DoublyLinkedList) ForwardList() (gather []int) {
	current := dll.head
	for current != nil {
		gather = append(gather, current.data)
		current = current.next
	}
	return
}

// BackwardList prints the elements of the list from tail to head.
// (连 尾端 开始收集资料)
func (dll *DoublyLinkedList) BackwardList() (gather []int) {
	current := dll.tail
	for current != nil {
		gather = append(gather, current.data)
		current = current.prev
	}
	return
}
