package trieTree

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	estimateApiScale = 10
	separator        = "/"
)

// trieTree defines the trieTree structure
type trieTree struct {
	// A better name like child or childNode would be more clear
	child  map[string]*trieTree
	status trieStatus
}

// trieStatus defines the status of the trieTree
type trieStatus struct {
	passCount int
	isEnd     bool
}

/*
AddPath defines a trieTree data structure and implements an AddPath method to add a path to the trie.
The trieTree structure uses a map to store child nodes, enabling fast lookups and efficient storage of many paths with shared prefixes.
*/
func (tree *trieTree) AddPath(path string) (err error) {
	paths := strings.Split(path, separator)

	var next *trieTree
	next = tree
	for i := 0; i < len(paths); i++ {
		if paths[i] != "" {
			// initialize maps if they are nil
			if next == nil {
				next = &trieTree{
					child: make(map[string]*trieTree, estimateApiScale),
				}
			}
			if next.child == nil {
				next.child = make(map[string]*trieTree, estimateApiScale)
			}

			// Add the path to the trieTree
			if _, ok := next.child[paths[i]]; !ok {
				next.child[paths[i]] = &trieTree{}
			}

			// alter the status of the trieTree
			next.child[paths[i]].status.passCount++

			if i == len(paths)-1 {
				next.child[paths[i]].status.isEnd = true
			}

			// move to the next node
			next = next.child[paths[i]]
		}
	}

	return
}

/*
Tree prints the trieTree structure in a tree-like format.
Using tab to control the indentation of the output. (用 tab 去控制缩排的输出)
*/
func (tree *trieTree) Tree(tab ...int) {
	// return if the current node has no child nodes to avoid to print the unnecessary tab
	if len(tree.child) == 0 {
		// Return to the parent node
		return
	}
	/*
		If tab is empty, initialize it to 0
		It means the first level of the tree (如果 tab 是空的，初始化它为 0，代表树的第一层)
	*/
	if len(tab) <= 0 {
		tab = []int{0}
	}

	// Parent node increase the tab value by 1 and then pass it to the child node.
	// (父节点增加 tab 的值，然后传递给子节点)
	tab[0]++
	// Print all the child nodes
	for key, value := range tree.child {
		for i := 0; i < tab[0]; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("|_ %s (Pass: %d, End: %t) \n", key, value.status.passCount, value.status.isEnd)
		// Pass the tab value to the child node !
		value.Tree(tab[0])
	}
}

// Test_Check_trieTree tests printing the trieTree structure in a tree-like format.
func Test_Check_trieTree(t *testing.T) {
	// After adding the path, the trieTree should be printed in a tree-like format
	t.Run("test printing trieTree", func(t *testing.T) {
		// Prepare the new stdout in order to capture the output of the function
		oldStdout := os.Stdout
		readFromNewStdout, newStdout, err := os.Pipe()
		require.NoError(t, err)
		os.Stdout = newStdout

		// Initialize the trieTree
		var tree = &trieTree{
			child: make(map[string]*trieTree, estimateApiScale),
		}
		// Add path to the trieTree
		err = tree.AddPath("/path1/path2/path3")
		require.NoError(t, err)
		err = tree.AddPath("/path1/path2/path3/path4")
		require.NoError(t, err)
		err = tree.AddPath("/path1/path2/path5")
		require.NoError(t, err)
		err = tree.AddPath("/path1/path2/path5/path6")
		require.NoError(t, err)
		err = tree.AddPath("/path1/path2/path5/path7")
		require.NoError(t, err)
		// Print the trieTree
		tree.Tree()

		// Count the number of each path, used to check the output later
		// (统计每个路径的数量，之后用于后面检查输出 !)
		countPath1 := "5"
		countPath2 := "5"
		countPath3 := "2"
		countPath4 := "1"
		countPath5 := "3"
		countPath6 := "1"
		countPath7 := "1"

		// Check if the path is the end of the path, used to check the output later
		// (检查路径是否是路径的结尾，之后用于后面检查输出 !)
		isEndPath1 := "false"
		isEndPath2 := "false"
		isEndPath3 := "true"
		isEndPath4 := "true"
		isEndPath5 := "true"
		isEndPath6 := "true"
		isEndPath7 := "true"

		// Close the new stdout and recover the original stdout
		_ = newStdout.Close()
		os.Stdout = oldStdout

		// Read the output of the function
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, readFromNewStdout)
		output := buf.String()

		// Check the output

		// Every line should be printed with certain tab correctly
		require.Equal(t, strings.Contains(output, "\t|_ path1"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t|_ path2"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t\t|_ path3"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t\t\t|_ path4"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t\t|_ path5"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t\t\t|_ path6"), true)
		require.Equal(t, strings.Contains(output, "\n\t\t\t\t|_ path7"), true)

		// The number of each path should be printed correctly
		require.Equal(t, strings.Contains(output, "path1 (Pass: "+countPath1+", End: "+isEndPath1), true)
		require.Equal(t, strings.Contains(output, "path2 (Pass: "+countPath2+", End: "+isEndPath2), true)
		require.Equal(t, strings.Contains(output, "path3 (Pass: "+countPath3+", End: "+isEndPath3), true)
		require.Equal(t, strings.Contains(output, "path4 (Pass: "+countPath4+", End: "+isEndPath4), true)
		require.Equal(t, strings.Contains(output, "path5 (Pass: "+countPath5+", End: "+isEndPath5), true)
		require.Equal(t, strings.Contains(output, "path6 (Pass: "+countPath6+", End: "+isEndPath6), true)
		require.Equal(t, strings.Contains(output, "path7 (Pass: "+countPath7+", End: "+isEndPath7), true)

		// Print the output finally
		fmt.Println(output)
	})
}
