package url

import (
	"errors"
	"strings"
)

// A URL represents a parsed URL
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses a raw URL to a URL struct
func Parse(rawurl string) (*URL, error) {

	scheme, rest, ok := strings.Cut(rawurl, ":")
	if !ok || scheme == "" {
		return nil, errors.New("missing scheme in URL")
	}

	if !strings.HasPrefix(rest, "//") {
		return &URL{Scheme: scheme}, nil
	}

	host, path, _ := strings.Cut(rest[2:], "/")

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
}

// String returns the string representation of the URL
func (u *URL) String() string {
	if u == nil {
		return ""
	}

	const lenSchemeSeparator = len("://")
	const lenPathSeparator = len("/")

	lenUrl := len(u.Scheme) + lenSchemeSeparator + len(u.Host) + lenPathSeparator + len(u.Path)

	var s strings.Builder
	s.Grow(lenUrl)

	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		if u.Host != "" {
			s.WriteString("://")
		} else {
			s.WriteString(":")
		}
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		if u.Host != "" {
			s.WriteString("/")
		}
		s.WriteString(p)
	}

	return s.String()
}
