# kmp
String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP)

See [Documentation](http://godoc.org/github.com/SimpCosm/algorithms/strings/kmp) on GoDoc.

## Example:

```go
package main

import (
  	"fmt"
	"github.com/paddie/gokmp"
)

const str = "aabaabaaaabbaabaabaaabbaabaabb"
const pattern = "aabb"

func main() {
	kmp, _ := gokmp.NewKMP(pattern)
	ints := kmp.FindAllStringIndex(str)

	fmt.Println(ints)
}

```

## Output:

```go
[8 19 26]
```



## Tests and Benchmarks:

```bash
$ go test -v kmp_test.go kmp.go -bench=.
```



## Output:

```go
=== RUN   TestFindAllStringIndex
--- PASS: TestFindAllStringIndex (0.00s)
=== RUN   TestFindStringIndex
--- PASS: TestFindStringIndex (0.00s)
=== RUN   TestOccurrences
--- PASS: TestOccurrences (0.00s)
=== RUN   TestOccurrencesFail
--- PASS: TestOccurrencesFail (0.00s)
=== RUN   TestContainedIn
--- PASS: TestContainedIn (0.00s)
goos: linux
goarch: amd64
BenchmarkKMPIndexComparison-8           100000000               14.9 ns/op
BenchmarkStringsIndexComparison-8       200000000                6.43 ns/op
PASS
ok      command-line-arguments  3.451s
```



## Comparison:

