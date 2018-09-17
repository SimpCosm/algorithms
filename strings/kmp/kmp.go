package kmp

import (
	"errors"
	"fmt"
)

type KMP struct {
	pattern string
	prefix  []int
	size    int
}

func (kmp *KMP) String() string {
	return fmt.Sprintf("pattern: %v\nprefix: %v", kmp.pattern, kmp.prefix)
}

// returns an array containing indexes of matches
// - error if pattern argument is less than 1 char
func buildTable(pattern string) ([]int, error) {

}

// compile new prefix-arrary given argument
func NewKMP(pattern string) (*KMP, error) {
	prefix, err := buildTable(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
		pattern: pattern,
		prefix:  prefix,
		size:    len(pattern),
	}, nil
}

// return index of first occurence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {

}

// find every occcurence of the kmp.pattern in 's'
func (kmp *KMP) FindAllStringIndex(s string) []int {

}

// returns true if pattern i matched at least once
func (kmp *KMP) ContainedIn(s string) bool {
	return kmp.FindStringIndex(s) >= 0
}

// returns the number of occurences of pattern in argument
func (kmp *KMP) Occurrences(s string) int {
	return len(kmp.FindAllStringIndex(s))
}
