// Package kmp implements the Knuth-Morris-Pratt algorithm (KMP) in Golang.
//
// The KMP algorithm searches for occurrences of a `word` W within a main `text string` S
// by employing the observation that when mismatch occurs, the word itself embodies sufficient
// information to determine where the next match could begin, the bypassing re-examination of
// previously matched characters.
package kmp

import (
	"errors"
	"fmt"
)

type KMP struct {
	pattern string
	next    []int
	size    int
}

func (kmp *KMP) String() string {
	return fmt.Sprintf("pattern: %v\nnext: %v", kmp.pattern, kmp.next)
}

// returns an array containing indexes of matches
// - error if pattern argument is less than 1 char
func buildTable(pattern string) ([]int, error) {
	lenP := len(pattern)
	if lenP == 0 {
		return nil, errors.New("'pattern' must contain at least one character")
	}

	i, j := 0, -1
	t := make([]int, lenP)
	t[0] = -1

	for i < lenP-1 {
		// pattern[j] is prefix, pattern[i] is suffix
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			t[i] = j
		} else {
			j = t[j]
		}
	}
	return t, nil
}

// NewKMP compile new next-table from given pattern.
func NewKMP(pattern string) (*KMP, error) {
	next, err := buildTable(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
		pattern: pattern,
		next:    next,
		size:    len(pattern),
	}, nil
}

// FindStringIndex returns index of first occurence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {
	// sanity check
	if len(s) < kmp.size {
		return -1
	}

	j, k := 0, 0
	for j < len(s) {
		if kmp.pattern[k] == s[j] {
			j++
			k++
			if k == kmp.size {
				return j - k
			}
		} else {
			k = kmp.next[k]
			if k < 0 {
				j++
				k++
			}
		}
	}
	return -1
}

// FindAllStringIndex finds every occcurence of the kmp.pattern in 's'
func (kmp *KMP) FindAllStringIndex(s string) []int {
	lenS := len(s)

	if lenS < kmp.size {
		return []int{}
	}

	indice := make([]int, 0, 10)

	j, k := 0, 0
	for j < len(s) {
		if kmp.pattern[k] == s[j] {
			j++
			k++
			if k == kmp.size {
				indice = append(indice, j-k)
				k = kmp.next[k-1] + 1
			}
		} else {
			k = kmp.next[k]
			if k < 0 {
				j++
				k++
			}
		}
	}
	return indice
}

// ContainedIn returns true if pattern i matched at least once
func (kmp *KMP) ContainedIn(s string) bool {
	return kmp.FindStringIndex(s) >= 0
}

// Occurrences returns the number of occurences of pattern in argument
func (kmp *KMP) Occurrences(s string) int {
	return len(kmp.FindAllStringIndex(s))
}
