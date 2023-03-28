package maxSubArrayLen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// maxSubArrayLen uses a hash map to efficiently find the maximum subarray sum of a given array with a target sum value.
func maxSubArrayLen(nums []int, targetSum int) (begin, end, length int) {
	/*
		Initialize a hash map to store the sum from the beginning of the array to the current position
		The initial value is {0:-1}, indicating that the sum from the beginning of the array to the -1st position is sum 0
	*/
	sumMap := map[int]int{
		0: -1,
	}

	// Initialize sum to 0, indicating that the sum from the beginning of the array to the -1st position is 0
	sum := 0
	for key, value := range nums {
		// Calculate the sum from the beginning of the array to the current position
		sum += value
		// Store the sum and its position into the sumMap
		sumMap[sum] = key
		// Check if there exists a previous sum equals to sum-targetSum in the sumMap
		if previous, ok := sumMap[sum-targetSum]; ok {
			// If the difference between the current position and the previous position is greater than the previous maximum length,
			// update the maximum length and the beginning and ending positions of the maximum subarray
			if key-previous > length {
				length = key - previous
				begin = previous
				end = key
			}
		}
	}

	// Return the beginning position, ending position, and length of the maximum subarray
	return
}

/*
Test_Check_maxSubArrayLen checks the maxSubArrayLen function with two examples.
The test ensures that the expected subarray length, beginning index, and ending index are returned for each example.
*/
func Test_Check_maxSubArrayLen(t *testing.T) {
	// The first example uses a simple array of integers to calculate the maximum subarray length with a given sum
	t.Run("Test for the maxSubArrayLen function", func(t *testing.T) {
		// Define an array of integers and calculate the maximum subarray length with a given sum
		arr := []int{1, 2, 3, 4, 5, 6}
		begin, end, length := maxSubArrayLen(arr, 12)

		// Check that the expected subarray length, beginning index and ending index are returned
		require.Equal(t, 3, length)
		// The sum should be calculated from the next index in the array
		require.Equal(t, 2, begin+1)
		require.Equal(t, 4, end)
	})
	// The second example uses a more complicated array of integers to calculate the same
	t.Run("Use a more complicated array", func(t *testing.T) {
		// The second example uses a more complicated array of integers to calculate the same
		arr := []int{5, 17, -1, 9, -4, 3, 10, 15, -3, 20, 16, 0, 19, 11, -5, 14, 12, -2, 7, 6}
		begin, end, length := maxSubArrayLen(arr, 138)

		// Check that the expected subarray length, beginning index and ending index are returned
		require.Equal(t, 18, length)
		// The sum should be calculated from the next index in the array
		require.Equal(t, 1, begin+1)
		require.Equal(t, 18, end)
	})
}
