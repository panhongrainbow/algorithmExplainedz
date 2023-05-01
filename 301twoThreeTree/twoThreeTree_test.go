package twoThreeTree

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

// TwoThreeTree defines the trieTree TwoThreeTree
type TwoThreeTree struct {
	// parent node
	parent *TwoThreeTree
	// keep three values
	middleNext []int
	// hold 3 next nodes
	leftNext *TwoThreeTree
	// middleNext moved
	rightNext *TwoThreeTree
}

// Create a new sub TwoThreeTree
func NewTwoThreeTree(parent *TwoThreeTree) (tree *TwoThreeTree) {
	// Initialize the value of the node
	tree = new(TwoThreeTree)
	// link to parent node
	tree.parent = parent
	// make the capacity of middle node
	tree.middleNext = make([]int, 0, 10)
	// Initialize the next node
	tree.leftNext = new(TwoThreeTree)
	tree.rightNext = new(TwoThreeTree)
	// Return the node
	return
}

// Insert a value into the TwoThreeTree
func (tree *TwoThreeTree) InsertValue(value int) (returnValue int, err error) {
	// Initialize the next nodes
	if tree.leftNext == nil {
		tree.leftNext = NewTwoThreeTree(tree)
	}
	if tree.rightNext == nil {
		tree.rightNext = NewTwoThreeTree(tree)
	}

	// If the value is 0, return
	if value == 0 {
		return
	}

	// Check the direction of the value
	direction := tree.Direction(value)

	// If the direction is left, insert the value into the left node
	if direction == left {
		returnValue, err = tree.leftNext.InsertValue(value)
		if err != nil {
			return
		}
		if returnValue == 0 {
			return
		}
		return
	}

	// If the direction is right, insert the value into the right node
	if direction == right {
		returnValue, err = tree.rightNext.InsertValue(value)
		if err != nil {
			return
		}
		if returnValue == 0 {
			return
		}
	}

LOOP:
	// If the direction is middle, after the middle node is full, insert the value into the parent node
	next := &TwoThreeTree{}
	if direction == middle {
		tree.middleNext = append(tree.middleNext, value)
		sort.Ints(tree.middleNext)
		if len(tree.middleNext) >= 10 {
			if tree.parent.middleNext[len(tree.parent.middleNext)-1] < tree.middleNext[0] {
				returnValue = tree.middleNext[0]
				tree.middleNext = tree.middleNext[1:]
				next = tree.parent
			} else {
				returnValue = tree.middleNext[len(tree.middleNext)-1]
				tree.middleNext = tree.middleNext[:len(tree.middleNext)-1]
				next = tree.rightNext
			}
		}
	}

	// If the return value in not 0, recursively insert the value into the parent node
	if returnValue != 0 {
		returnValue, err = next.InsertValue(returnValue)
		if err != nil {
			return
		}
		if returnValue != 0 {
			goto LOOP
		}
	}

	// return
	return
}

const (
	left = iota
	middle
	right
)

func (tree *TwoThreeTree) Direction(value int) (direction int) {
	if len(tree.middleNext) < 2 {
		direction = middle
		return
	}
	if tree.middleNext[0] > value {
		direction = left
		return
	}
	if tree.middleNext[len(tree.middleNext)-1] < value {
		direction = right
		return
	}
	direction = middle
	return
}

func Test_Check_twoThreeTree(t *testing.T) {
	// Create a new TwoThreeTree
	root := NewTwoThreeTree(nil)
	var rootReturnValue int
	var err error

	rootReturnValue, err = root.InsertValue(25)
	require.NoError(t, err)
	require.Equal(t, 0, rootReturnValue)

	rootReturnValue, err = root.InsertValue(50)
	require.NoError(t, err)
	require.Equal(t, 0, rootReturnValue)

	rootReturnValue, err = root.InsertValue(75)
	require.NoError(t, err)
	require.Equal(t, 0, rootReturnValue)

	rootReturnValue, err = root.InsertValue(100)
	require.NoError(t, err)
	require.Equal(t, 0, rootReturnValue)

	for i := 76; i < 83; i++ {
		rootReturnValue, err = root.InsertValue(i)
		require.NoError(t, err)
		require.Equal(t, 0, rootReturnValue)
	}

	rootReturnValue, err = root.InsertValue(84)
	require.NoError(t, err)
	require.Equal(t, 0, rootReturnValue)

}
