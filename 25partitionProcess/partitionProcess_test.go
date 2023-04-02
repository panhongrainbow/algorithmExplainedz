package partitionProcess

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
partitionProcess divides an array into three sections based on a pivot value.
Elements less than, equal to, and greater than the pivot value
are placed in the leftmost, middle, and rightmost sections respectively.
*/
func partitionProcess(arr []int, pivotValue int) (err error) {
	/*
		The array will be divided into three sections, with
		the leftmost section consisting of elements in the array that are all less than the pivot value,
		the middle section consisting of elements that are all equal to the pivot value,
		and
		the rightmost section consisting of elements that are all greater than the pivot value.
	*/
	// the boundaries of the middle section
	var middleSectionRightBoundary = 0
	var middleSectionLeftBoundary = 0
	// the left boundary of the right most section
	var rightMostSectionLeftBoundary = len(arr) - 1

	// Iterate over each element in the array
	for i := 0; i < len(arr); i++ {
		if arr[i] == pivotValue {
			/*
				If the current element is equal to the pivot value,
				move it to the right within the boundary of the middle section by one position
			*/
			middleSectionRightBoundary++
			// Swap the current element with the last second element to the right of the middle section
			swapArr(arr, i, middleSectionRightBoundary-1)
		}
		if arr[i] < pivotValue {
			/*
				If the current element is less than the pivot value,
				move it to the right within the boundary of the middle section by one position
			*/
			middleSectionRightBoundary++
			/*
				If the current element is less than the pivot value,
				move it to the left within the boundary of the middle section by one position
			*/
			middleSectionLeftBoundary++
			// Swap the current element with the previous element to the left of the middle section
			swapArr(arr, i, middleSectionLeftBoundary-1)
		}
		// If the current element is greater than the pivot value,
		if arr[i] > pivotValue {
			// Swap the current element with the first element to the left of the rightmost section
			swapArr(arr, i, rightMostSectionLeftBoundary)

			// Move the left boundary of the rightmost section to the left by one position
			rightMostSectionLeftBoundary--

			// i index doesn't move forward
			if i > 0 {
				i--
			}
		}
		/*
			If the boundaries of the middle section and the rightmost section intersect,
			it means that all elements have been processed and the loop can be terminated.
		*/
		/*
			This statement needs to be placed here; otherwise, the loop may not stop immediately, leading to an error.
			必须放在这里，否则循环不会及时停止，会发生错误。
			必須放在這，都已經正確了，如果再排下去會出錯
		*/
		if middleSectionRightBoundary > rightMostSectionLeftBoundary {
			break
		}
	}
	// Return nil to indicate that the function executed successfully
	return nil
}

// swapArr swaps two elements of an integer array using a temporary variable.
func swapArr(arr []int, i, j int) {
	a := arr[i]
	b := arr[j]
	arr[i] = b
	arr[j] = a
	return
}

/*
Test_Check_partitionProcess is a test function that tests the "partitionProcess" function using different arrays of integers.
Check whether elements that are less than, equal to, and greater than the pivot value
are placed in the leftmost, middle, and rightmost sections, respectively.
*/
func Test_Check_partitionProcess(t *testing.T) {
	// Tests the partitionProcess function by using a simple example
	t.Run("Test case uses a simple array", func(t *testing.T) {
		// Define an input array of integers
		arr := []int{3, 5, 2, 6, 4, 1}

		// Call the partitionProcess function and ensure no error is returned
		err := partitionProcess(arr, 6)
		require.NoError(t, err)

		// Check that the first five elements of the array are less than 6
		for i := 0; i < len(arr)-1; i++ {
			require.Greater(t, 6, arr[i])
		}
		// Check that the element at the last partition index is equal to 6
		require.Equal(t, 6, arr[len(arr)-1])

		// fmt.Println(arr)
		// The output is [3 5 2 4 1 6]
	})
	t.Run("Test case uses a more complicated array", func(t *testing.T) {
		// Define an input array of integers
		arr := []int{33, 47, 95, 47, 31, 56, 76, 47, 16, 63, 47, 90, 24}

		// Call the partitionProcess function and ensure no error is returned
		err := partitionProcess(arr, 47)
		require.NoError(t, err)

		// Check that the first four elements of the array are less than 3
		for i := 0; i < 4; i++ {
			require.Greater(t, 47, arr[i])
		}
		// Check if the elements from the fifth to the eighth are equal to 47
		for i := 4; i < 8; i++ {
			require.Equal(t, 47, arr[i])
		}
		// Check that the last five elements of the array are greater than 3
		for i := 8; i < len(arr); i++ {
			require.Less(t, 47, arr[i])
		}

		// fmt.Println(arr)
		// The output is [33 24 31 16 47 47 47 47 63 76 90 56 95]
	})
	// The second example uses a more complicated array of integers to calculate the same
	t.Run("provide another more complicated array", func(t *testing.T) {
		// Define a more complicated input array of integers
		arr := []int{3, 3, 3, 3, 3, 5, 17, 3, -1, 9, -4, 3, 10, 15, -3, 20, 3, 3, 3, 3, 16, 0, 19, 11, -5, 14, 12, -2, 7, 6, 3, 3, 3, 3}

		// Call the partitionProcess function and ensure no error is returned
		err := partitionProcess(arr, 3)
		require.NoError(t, err)

		// Check that the first sixth elements of the array are less than 3
		for i := 0; i < 6; i++ {
			require.Greater(t, 3, arr[i])
		}
		// Check if the elements from the seventh to the Twenty-first are equal to 3
		for i := 6; i < 21; i++ {
			require.Equal(t, 3, arr[i])
		}
		// Check that the last thirteen elements of the array are greater than 3
		for i := 21; i < len(arr); i++ {
			require.Less(t, 3, arr[i])
		}

		// fmt.Println(arr)
		// The output is [-1 -4 -2 -3 -5 0 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3 19 11 16 14 12 20 7 6 15 10 9 17 5]
	})
}
