package ubi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	var testCases []func()

	b0 := NewBuilder(url.URL{})

	b1 := b0.SetScheme("https")
	testCases = append(testCases, func() { assert.Equal(t, "https:", b1.URL().String()) })
	doAll(testCases)

	b2 := b1.SetHost("example.com:8080")
	testCases = append(testCases, func() { assert.Equal(t, "https://example.com:8080", b2.URL().String()) })
	doAll(testCases)

	b3 := b2.SetPaths("foo")
	testCases = append(testCases, func() { assert.Equal(t, "https://example.com:8080/foo", b3.URL().String()) })
	doAll(testCases)

	b4 := b3.AppendPaths("bar", "buz")
	testCases = append(testCases, func() { assert.Equal(t, "https://example.com:8080/foo/bar/buz", b4.URL().String()) })
	doAll(testCases)

	b5 := b4.SetPaths("qux", "foobar")
	testCases = append(testCases, func() { assert.Equal(t, "https://example.com:8080/qux/foobar", b5.URL().String()) })
	doAll(testCases)

	b6 := b5.SetQuery(url.Values{
		"q1": []string{"v1"},
	})
	testCases = append(testCases, func() { assert.Equal(t, "https://example.com:8080/qux/foobar?q1=v1", b6.URL().String()) })
	doAll(testCases)

	b7 := b6.AddQuery(url.Values{
		"q2": []string{"v2"},
		"q3": []string{"v3-1", "v3-2"},
	})
	testCases = append(testCases, func() {
		assert.Equal(t, "https://example.com:8080/qux/foobar?q1=v1&q2=v2&q3=v3-1&q3=v3-2", b7.URL().String())
	})
	doAll(testCases)

	b8 := b7.SetQuery(url.Values{
		"q4": []string{"v4"},
		"q5": []string{"v5-1", "v5-2"},
	})
	testCases = append(testCases, func() {
		assert.Equal(t, "https://example.com:8080/qux/foobar?q4=v4&q5=v5-1&q5=v5-2", b8.URL().String())
	})
	doAll(testCases)

	b9 := b8.SetFragment("frag")
	testCases = append(testCases, func() {
		assert.Equal(t, "https://example.com:8080/qux/foobar?q4=v4&q5=v5-1&q5=v5-2#frag", b9.URL().String())
	})
	doAll(testCases)
}

func TestBuilder_SetHostWithEscaping(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	b := NewBuilder(*u)
	assert.Equal(t, "https://example%2F.com", b.SetHost("example/.com").URL().String())
}

func TestBuilder_SetPathsWithEscaping(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	b := NewBuilder(*u)
	assert.Equal(t, "https://example.com/foo%252Fbar/buz%253Fqux", b.SetPaths("foo/bar", "buz?qux").URL().String())
}

func TestBuilder_AppendPathsWithEscaping(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	b := NewBuilder(*u)
	assert.Equal(t, "https://example.com/foo%252Fbar/buz%253Fqux", b.AppendPaths("foo/bar", "buz?qux").URL().String())
}

func TestBuilder_SetQueryWithEscaping(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	b := NewBuilder(*u)
	assert.Equal(t, "https://example.com?foo%3Fbar=buz%26qux", b.SetQuery(url.Values{"foo?bar": []string{"buz&qux"}}).URL().String())
}

func TestBuilder_AddQueryWithEscaping(t *testing.T) {
	u, _ := url.Parse("https://example.com")
	b := NewBuilder(*u)
	assert.Equal(t, "https://example.com?foo%3Fbar=buz%26qux", b.AddQuery(url.Values{"foo?bar": []string{"buz&qux"}}).URL().String())
}

func doAll(testCases []func()) {
	for _, testCase := range testCases {
		testCase()
	}
}
