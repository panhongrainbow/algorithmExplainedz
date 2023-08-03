package MaxHeap

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// siftUp is sift up operation (大根堆的上移操作)
func siftUp(heap []int, idx int) {
	// The index "idx" cannot be 0 because index 0 does not have a parent node
	for idx > 0 {
		// Find the index of the parent node and truncating the decimal point
		parentIdx := (idx - 1) / 2
		// In a max heap, if the current node is greater than its parent node, they should be swapped.
		if heap[parentIdx] < heap[idx] {
			heap[parentIdx], heap[idx] = heap[idx], heap[parentIdx]
			idx = parentIdx
		} else {
			// Break the loop
			break
		}
	}
}

// siftDown is sift up operation (大根堆的下移操作)
func siftDown(heap []int, idx, heapSize int) {
	// The index must be within the heap size.
	for idx < heapSize {
		// Find the indices of the left and right child nodes, denoted as leftChildIdx and rightChildIdx, respectively
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2

		// First, we need to determine whether the left child node is larger or the right child node is larger.
		// This is because only the largest value can reach the first element's position, which is the position of the root node in the max heap.
		// (只有最大的值能到第一个元素，也就是大根堆的根结点的位置
		var rightNodeBigger bool
		if leftChildIdx < heapSize && rightChildIdx < heapSize {
			if heap[rightChildIdx] > heap[leftChildIdx] {
				rightNodeBigger = true
			}
		}

		// If the parent node is smaller than the left child node, swap them.
		// Before performing the comparison and swap, we should check if the indices of the left and right child nodes are within the valid range,
		// so we need to check if leftChildIdx or rightChildIdx is less than heapSize.
		// (就是要检查左右子节点是否有效，就检查是 index 是否小于 heap size 里面)
		var moved bool
		if rightNodeBigger == true {
			// If the right child node is larger, swap positions with the right child node.
			// (右边子节点比较大，就下移到子结点的位置，那里大就去那里)
			if rightChildIdx < heapSize &&
				heap[idx] < heap[rightChildIdx] {

				heap[rightChildIdx], heap[idx] = heap[idx], heap[rightChildIdx]
				idx = rightChildIdx

				moved = true
			}
		} else {
			// If the left child node is larger, swap positions with the left child node.
			if leftChildIdx < heapSize &&
				heap[idx] < heap[leftChildIdx] {

				heap[leftChildIdx], heap[idx] = heap[idx], heap[leftChildIdx]
				idx = leftChildIdx

				moved = true
			}
		}

		// // If both the left and right child nodes cannot be swapped, then break the loop;
		// otherwise, it may result in an infinite loop.
		// (如果左右子结点都无法换，就不要换了)
		if moved == false {
			break
		}
	}
}

// insert is a function used to add new data to the max heap
func insert(heap []int, value int) []int {
	heap = append(heap, value)
	// Perform the sift up operation on the newly added element at the end.
	// (对新增到最后面的元素进行上移操作)
	siftUp(heap, len(heap)-1)
	return heap
}

// extractMax is used to retrieve the maximum element from the max heap, which is also the first element in the slice.
func extractMax(heap []int) (int, []int) {
	if len(heap) == 0 {
		return 0, heap
	}

	// Reorganize the slice
	max := heap[0]
	lastIdx := len(heap) - 1
	// Move the last element to the position of index 0
	// (移动第一个元素到第一个位置，进行下移操作)
	heap[0] = heap[lastIdx]
	heap = heap[:lastIdx]
	siftDown(heap, 0, len(heap))

	return max, heap
}

// Test_Check_Max_Heap tests the insertion and extraction operations on a max heap more than 5000 times and
// verifies that the maximum values extracted from the max heap match those obtained from the sorted slice.
func Test_Check_Max_Heap(t *testing.T) {
	// >>>>> >>>>> >>>>> initial 初始化

	// Create an empty slice maxHeap to hold elements in max heap (大根堆)
	maxHeap := make([]int, 0, 20)

	// Create an empty slice maxSlice to hold elements to be sorted in descending order (就一般的切片)
	maxSlice := make([]int, 0, 20)

	// Create a local random generator using the current Unix time as the seed
	localRand := rand.New(rand.NewSource(time.Now().Unix()))

	// >>>>> >>>>> >>>>> insert data 开始塞资料

	// Generate and insert 20 random numbers into both maxSlice and maxHeap
	for i := 0; i < 20; i++ {

		// Generate a random number between 1 and 5000
		number := localRand.Intn(5000) + 1

		// Append the generated number to maxSlice and sort maxSlice in descending order
		maxSlice = append(maxSlice, number)
		sort.Slice(maxSlice, func(i, j int) bool {
			return maxSlice[i] > maxSlice[j]
		})

		// Insert the generated number into maxHeap
		maxHeap = insert(maxHeap, number)
	}

	// >>>>> >>>>> >>>>> insert and extract data 一边塞资料一边取最大值

	// Continue inserting 5000 more random numbers into both maxSlice and maxHeap,
	// and compare the extracted maximum values from maxSlice and maxHeap
	for i := 0; i < 5000; i++ {
		// Generate a random number between 1 and 5000
		randomNumber := localRand.Intn(5000) + 1

		// Append the generated number to maxSlice and sort maxSlice in descending order
		maxSlice = append(maxSlice, randomNumber)
		sort.Slice(maxSlice, func(i, j int) bool {
			return maxSlice[i] > maxSlice[j]
		})

		// Insert the generated number into maxHeap
		maxHeap = insert(maxHeap, randomNumber)

		// Extract the maximum value from maxSlice
		maxValueInSlice := maxSlice[0]
		maxSlice = maxSlice[1:]

		// Extract the maximum value from maxHeap
		var maxValueInMaxHeap int
		maxValueInMaxHeap, maxHeap = extractMax(maxHeap)

		// Assert that the maximum values from maxSlice and maxHeap are equal
		require.Equal(t, maxValueInSlice, maxValueInMaxHeap)
	}

	// >>>>> >>>>> >>>>> continue to extract for another 20 times 把大根堆和切片的剩下的 20 笔资料全部取出

	// After inserting 5000 random numbers, continue to extract and compare the maximum values
	// from maxSlice and maxHeap for another 20 times
	for i := 0; i < 20; i++ {
		// Extract the maximum value from maxSlice
		maxInSlice := maxSlice[0]
		maxSlice = maxSlice[1:]

		// Extract the maximum value from maxHeap
		var maxInMaxHeap int
		maxInMaxHeap, maxHeap = extractMax(maxHeap)

		// Assert that the maximum values from maxSlice and maxHeap are equal
		require.Equal(t, maxInSlice, maxInMaxHeap)
	}
}

// Benchmark_Benchmark_MaxHeap inserts and extracts data, testing the max heap (大根堆) performance
func Benchmark_Benchmark_MaxHeap(b *testing.B) {
	// >>>>> >>>>> >>>>> initial 初始化

	// Create an empty slice maxHeap to hold elements in max heap (大根堆)
	maxHeap := make([]int, 0, 200)

	// Create a local random generator using the current Unix time as the seed
	localRand := rand.New(rand.NewSource(time.Now().Unix()))

	// Reset benchmark timer
	b.ResetTimer()

	// >>>>> >>>>> >>>>> insert and extract data 一边塞资料一边取最大值

	// Generate and insert 20 random numbers into both maxSlice and maxHeap
	for i := 0; i < b.N; i++ {

		// Generate a random number between 1 and 5000
		number := localRand.Intn(5000) + 1

		// Insert the generated number into maxHeap
		maxHeap = insert(maxHeap, number)

		if i > 200 {
			// Extract the maximum value from maxHeap
			_, maxHeap = extractMax(maxHeap)
		}
	}
}

// Benchmark_Benchmark_MaxSlice inserts and extracts data, testing the sorted slice performance
func Benchmark_Benchmark_MaxSlice(b *testing.B) {
	// >>>>> >>>>> >>>>> initial 初始化

	// Create an empty slice maxSlice to hold elements to be sorted in descending order (就一般的切片)
	maxSlice := make([]int, 0, 20)

	// Create a local random generator using the current Unix time as the seed
	localRand := rand.New(rand.NewSource(time.Now().Unix()))

	// Reset benchmark timer
	b.ResetTimer()

	// >>>>> >>>>> >>>>> insert and extract data 一边塞资料一边取最大值

	// Continue inserting 5000 more random numbers into both maxSlice and maxHeap,
	// and compare the extracted maximum values from maxSlice and maxHeap
	for i := 0; i < b.N; i++ {
		// Generate a random number between 1 and 5000
		randomNumber := localRand.Intn(5000) + 1

		// Append the generated number to maxSlice and sort maxSlice in descending order
		maxSlice = append(maxSlice, randomNumber)
		sort.Slice(maxSlice, func(i, j int) bool {
			return maxSlice[i] > maxSlice[j]
		})

		if i > 200 {
			// Extract the maximum value from maxSlice
			maxSlice = maxSlice[1:]
		}
	}
}
