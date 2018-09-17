package trie

import (
    "sort"
)

type Node struct {
    val         rune
    isEnd       bool
    children    map[rune]*Node
}

type Trie struct {
    root    *Node
    size    int
}

// Creates a new Trie with an initialized root Node
func New() *Trie {
    return &Trie{
        root: &Node{children: make(map[rune]*Node), depth: 0},
        size: 0,
    }
}

// Returns the root node for the Trie
func (t *Trie) Root() *Node {
    return t.root
}

// Adds the key to the Trie, including meta data.
// Meta data is stored as `interface{}` and must be type cast by the caller.
func (t *Trie) Add(key string, meta interface{}) *Node {

}

// Removes a key from the trie, ensuring that
// all bitmasks up to root are approriately recalculated.
func (t *Trie) Remove(key string) {
    
}

// Finds and returns meta data associated with `key`
func (t *Trie) Find(key string) (*Node, bool) {

}

func (t *Trie) HasKeysWithPrefix(key string) bool {

}

// Returns all the keys currently stored in the trie
func (t *Trie) Keys() []string {

}

// Return a fuzzy search against the keys in the trie.
func (t *Trie) FuzzySearch(pre string) []string {

}

// Perform a prefix search against the keys in the trie.
func (t *Trie) PrefixSearch(pre string) []string {

}
