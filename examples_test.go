package ubi

import (
	"fmt"
	"net/url"
)

func ExampleNewBuilder() {
	b := NewBuilder(&url.URL{}) // or the parameter can be nil
	u := b.SetScheme("https").
		SetHost("example.com:8080").
		SetPaths("foo", "bar").
		AppendPaths("buz", "qux").
		SetQuery(url.Values{"q1": []string{"v1"}, "q2": []string{"v2-1", "v2-2"}}).
		AddQuery(url.Values{"q3": []string{"v3"}}).
		SetFragment("frag").
		URL()
	fmt.Printf("%s\n", u)

	// example of immutability
	b0 := NewBuilder(&url.URL{})
	b1 := b0.SetScheme("https").
		SetHost("example.com")
	fmt.Printf("b1: %s\n", b1.URL())

	b2 := b1.SetHost("example.com:8080")
	fmt.Printf("b1 (immutable: previous one still remembers the state): %s\n", b1.URL())
	fmt.Printf("b2 (new one): %s\n", b2.URL())

	// Output:
	// https://example.com:8080/foo/bar/buz/qux?q1=v1&q2=v2-1&q2=v2-2&q3=v3#frag
	// b1: https://example.com
	// b1 (immutable: previous one still remembers the state): https://example.com
	// b2 (new one): https://example.com:8080
}
