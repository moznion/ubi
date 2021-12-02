package ubi

import (
	"net/url"
	"path"
)

// Builder is a URL builder struct.
// This builder behaves with immutability; this implies all of the methods return a new Builder instance.
// It would be useful because every phase remembers its own state.
type Builder struct {
	url *url.URL
}

// NewBuilder is a constructor of Builder.
// When it needs to build an URL from the empty, please give the default value or nil to this function.
func NewBuilder(u *url.URL) *Builder {
	return &Builder{
		url: func() *url.URL {
			if u == nil {
				return &url.URL{}
			}
			return u
		}(),
	}
}

// SetScheme sets an URL scheme (e.g. "https").
func (b *Builder) SetScheme(scheme string) *Builder {
	u := *b.url
	u.Scheme = scheme
	return NewBuilder(&u)
}

// SetHost sets the host (e.g. "example.com:8080").
// This method escapes the given string by path escaping.
func (b *Builder) SetHost(host string) *Builder {
	u := *b.url
	u.Host = host
	return NewBuilder(&u)
}

// SetPaths sets the URL paths.
// This *replaces* the paths with the original ones. If it needs to append the paths, please consider using AppendPaths().
// This method escapes the given paths by path escaping.
func (b *Builder) SetPaths(paths ...string) *Builder {
	return b.setPaths("", paths...)
}

// AppendPaths appends the URL paths.
// This *appends* the paths to the original ones. If it needs to replace the paths, please consider using SetPaths().
// This method escapes the given paths by path escaping.
func (b *Builder) AppendPaths(paths ...string) *Builder {
	return b.setPaths(b.url.Path, paths...)
}

// SetQuery sets the query parameters.
// This *replaces* the parameters to the original ones. If it needs to add the parameters, please consider using AddQuery().
// This method escapes the given parameters by query escaping.
func (b *Builder) SetQuery(queryParameter url.Values) *Builder {
	u := *b.url
	u.RawQuery = queryParameter.Encode()
	return NewBuilder(&u)
}

// AddQuery adds the query parameters.
// This *adds* the parameters to the original ones. If it needs to replace the parameters, please consider using SetQuery().
// This method escapes the given parameters by query escaping.
func (b *Builder) AddQuery(queryParameter url.Values) *Builder {
	u := *b.url
	q := u.Query()
	for key, values := range queryParameter {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	return NewBuilder(&u)
}

// SetFragment sets the URL fragment.
func (b *Builder) SetFragment(fragment string) *Builder {
	u := *b.url
	u.Fragment = fragment
	return NewBuilder(&u)
}

func (b *Builder) URL() *url.URL {
	return b.url
}

func (b *Builder) setPaths(originalPath string, paths ...string) *Builder {
	u := *b.url

	joined := originalPath
	for _, p := range paths {
		joined = path.Join(joined, url.PathEscape(p))
	}

	u.Path = joined
	return NewBuilder(&u)
}
