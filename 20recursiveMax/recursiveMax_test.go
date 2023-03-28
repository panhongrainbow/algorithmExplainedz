package recursiveMax

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Error string

const (
	elementLenghtFinishInAdvance int = 3
)

const (
	ErrCannotFinishInAdvance = Error("cannot finish recursiveMax in advance")
	ErrEmptySlice            = Error("error happens because empty slice")
)

func (e Error) Error() string {
	return string(e)
}

// recursiveMax recursively finds the maximum value in an array, dividing it into two parts if its length exceeds a certain length.
func recursiveMax(arr []int) (temporaryMaxNumLeft int, err error) {
	/*
		If the length of the input array is greater than the elementLenghtFinishInAdvance,
		divide it into two parts and recursively find the maximum value
	*/
	if len(arr) > elementLenghtFinishInAdvance {
		// Find the middle index of the array
		var mid = len(arr) / 2

		// Find the maximum value in the left half of the array recursively
		var temporaryMaxNumRight int
		temporaryMaxNumLeft, err = recursiveMax(arr[0:mid])
		if err != nil {
			return
		}

		// Find the maximum value in the right half of the array recursively
		temporaryMaxNumRight, err = recursiveMax(arr[mid:])
		if err != nil {
			return
		}

		// Compare the maximum values of the left and right halves and return the larger one
		if temporaryMaxNumRight > temporaryMaxNumLeft {
			temporaryMaxNumLeft = temporaryMaxNumRight
		}

		// If the length of the input array is less than or equal to the threshold, find the maximum value directly
	} else if len(arr) <= elementLenghtFinishInAdvance {
		// Find the maximum value directly
		temporaryMaxNumLeft, err = finishEarly(arr)
	}

	// Return the maximum value and error
	return
}

// finishEarly finds the maximum value in a given array, but returns an error if the array is empty or larger than a specified limit.
func finishEarly(arr []int) (maxNum int, err error) {
	// If the length of the input array is less than or equal to 0, return an error
	if len(arr) <= 0 {
		err = ErrEmptySlice
		return
	}

	// If the length of the input array is greater than the elementLenghtFinishInAdvance, return an error
	if len(arr) > elementLenghtFinishInAdvance {
		err = ErrCannotFinishInAdvance
		return
	}

	// Find the maximum value in the input array
	maxNum = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}

	// Return the maximum value and error
	return
}

/*
recursiveMax consists of two functions, recursiveMax and finishEarly, that find the maximum value in an integer array.
The code is tested using the Go testing package with two test cases.
*/
func Test_Check_recursiveMax(t *testing.T) {
	// Tests the recursiveMax function by using a simple example
	t.Run("Test case uses a simple array", func(t *testing.T) {
		// Define an input array of integers
		arr := []int{1, 2, 3, 4, 5, 6}

		// Call the recursiveMax function and ensure no error is returned
		maxNum, err := recursiveMax(arr)
		require.NoError(t, err)

		// Assert that the maximum value of the input array is 6
		require.Equal(t, 6, maxNum)
	})
	// The second example uses a more complicated array of integers to calculate the same
	t.Run("Use a more complicated array", func(t *testing.T) {
		// Define a more complicated input array of integers
		arr := []int{5, 17, -1, 9, -4, 3, 10, 15, -3, 20, 16, 0, 19, 11, -5, 14, 12, -2, 7, 6}

		// Call the recursiveMax function with the complicated array and ensure no error is returned
		maxNum, err := recursiveMax(arr)
		require.NoError(t, err)

		// Assert that the maximum value of the complicated array is 20
		require.Equal(t, 20, maxNum)
	})
}
