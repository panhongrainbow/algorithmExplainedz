package trieTree

import (
	"fmt"
	"strings"
	"testing"
)

var TrieEntry *trieTree

type trieTree struct {
	next   map[string]*trieTree
	status trieStatus
}

type trieStatus struct {
	pass int
	end  bool
}

/*
Add defines a trieTree data structure and implements an Add method to add a path to the trie.
The trieTree structure uses a map to store child nodes, enabling fast lookups and efficient storage of many paths with shared prefixes.
*/
func (tree *trieTree) Add(path string) (err error) {
	paths := strings.Split(path, "/")

	var next *trieTree
	next = tree
	for i := 0; i < len(paths); i++ {
		if paths[i] != "" {
			if next == nil {
				next = &trieTree{
					next: make(map[string]*trieTree, 10),
				}
			}

			if next.next == nil {
				next.next = make(map[string]*trieTree, 10)
			}

			if _, ok := next.next[paths[i]]; !ok {
				next.next[paths[i]] = &trieTree{}
			}
			next.status.pass++
			next = next.next[paths[i]]
		}
	}

	return
}

func (tree *trieTree) List(tab ...int) {
	if len(tab) == 0 {
		tab = []int{0}
	}
	for key, value := range tree.next {
		for i := 0; i < tab[0]; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\\_ %s\n", key)
		tab[0]++
		value.List(tab...)
	}
	return
}

func Test_Check_trieTree(t *testing.T) {
	var tree = &trieTree{
		next: make(map[string]*trieTree, 10),
	}
	tree.Add("/path1/path2/path3")
	fmt.Println(tree)
	tree.List()
}
