package doublyLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {
	// Create a new doubly linked list.
	list := DoublyLinkedList{}

	// Append elements to the list.
	list.Append(10)
	list.Append(20)
	list.Append(30)

	assert.Equal(t, []int{10, 20, 30}, list.ForwardList())

	assert.Equal(t, []int{30, 20, 10}, list.BackwardList())

	list.Reverse()

	assert.Equal(t, []int{30, 20, 10}, list.ForwardList())

	assert.Equal(t, []int{10, 20, 30}, list.BackwardList())
}
