# ubi [![.github/workflows/check.yml](https://github.com/moznion/ubi/actions/workflows/check.yml/badge.svg)](https://github.com/moznion/ubi/actions/workflows/check.yml) [![codecov](https://codecov.io/gh/moznion/ubi/branch/main/graph/badge.svg?token=8GG9ECIHF7)](https://codecov.io/gh/moznion/ubi)

URL Builder with Immutability for golang.

This URL builder behaves with immutability; this implies all of the methods return a new Builder instance.
It would be useful because every phase remembers its own state.

## Synopsis

Basic Usage:

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/moznion/ubi"
)

func main() {
	b := NewBuilder(&url.URL{}) // or this parameter be nil
	u := b.SetScheme("https").
		SetHost("example.com:8080").
		SetPaths("foo", "bar").
		AppendPaths("buz", "qux").
		SetQuery(url.Values{"q1": []string{"v1"}, "q2": []string{"v2-1", "v2-2"}}).
		AddQuery(url.Values{"q3": []string{"v3"}}).
		SetFragment("frag").
		URL()
	fmt.Printf("%s\n", u)

	// Output:
	// https://example.com:8080/foo/bar/buz/qux?q1=v1&q2=v2-1&q2=v2-2&q3=v3#frag
}
```

Example of immutability:

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/moznion/ubi"
)

func main() {
	b0 := NewBuilder(&url.URL{})
	b1 := b0.SetScheme("https").
		SetHost("example.com")
	fmt.Printf("b1: %s\n", b1.URL())

	b2 := b1.SetHost("example.com:8080")
	fmt.Printf("b1 (immutable: previous one still remembers the state): %s\n", b1.URL())
	fmt.Printf("b2 (new one): %s\n", b2.URL())

	// Output:
	// b1: https://example.com
	// b1 (immutable: previous one still remembers the state): https://example.com
	// b2 (new one): https://example.com:8080
}
```

Examples are [here](./examples_test.go).

## Documentations

[![GoDoc](https://godoc.org/github.com/moznion/ubi?status.svg)](https://godoc.org/github.com/moznion/ubi)

## Author

moznion (<moznion@mail.moznion.net>)

