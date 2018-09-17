package bm

import (
    "fmt"
)

type BM struct {
    pattern string
    badCharSkip [256]int
    goodSuffixSkip []int
}

func (bm *BM) String() string {
    return fmt.Sprintf("pattern: %v\n", bm.pattern)
}

func NewBM(pattern string) (*BM, error) {

}

// next returns the index in the text of the first occurrence of the pattern.
// If the pattern is not found, it returns -1
func (bm *BM) next(text string) int {

}

func longestCommonSuffix(a, b string) (i int) {
    var lenA = len(a)
    var lenB = len(b)
    for i = 0; i < lenA && i < lenB; i++ {
        if a[lenA-1-i] != b[lenB-1-i] {
            break
        }
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
