package kmp_test

import (
    "fmt"
    "github.com/SimpCosm/algorithms/strings/kmp"
)

func Example() {
    str := "aabaabaaaabbaabaabaaabbaabaabb"
    pattern := "aabb"

    k, _ := kmp.NewKMP(pattern)
    ints := kmp.FindAllStringIndex(str)

    fmt.Println(ints)

    // Output:
    // [8 19 26]
}
