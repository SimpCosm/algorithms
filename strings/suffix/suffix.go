package suffix

import (
    "bytes"
    "sort"
)

func Edge struct {
    label   []byte
    point   interface{}
}

type Leaf struct {
    originKey   []byte
    value       interface{}
}

func Node struct {
    edges   []*Edge
}

// Tree represents a suffix tree.
type Tree struct {
    root    *Node
    leavesNum   int
}

// NewTree creates a suffix tree
func NewTree() *Tree {
    return &Tree{
        root: &Node {
            edges:  []&Edge{},
        },
        leavesNum: 0,
    }
}

// Insert suffix with given key and value.
// Return the previous value and a boolean to indicate whether the inseartion is succcessful.
func (t *Tree) Insert(key []byte, value interface{}) (oldValue interface{}, ok bool) {

}

// Get returns the value of given key and a boolean to indicate
// whether the value is found.
func (t *Tree) Get(key []byte) (value interface{}, found bool) {

}

// Remove returns the value of given key and a boolean to indicate
// whether the value is found. Then the value will be removed.
func (t *Tree) Remove(key []byte) (oldValue interface{}, found bool) {

}

// Walk through the tree, call function with key and value.
// Once the function returns true, it will stop walking.
// The travelling order is DFS, in the same suffix level the shortest key comes first.
func (t *Tree) Walk(f func(key []byte, value interface{}) bool) {
    stop := false
    t.root.walk([]byte{}, f, &stop)
}
