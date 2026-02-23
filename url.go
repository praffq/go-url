package url

import (
	"errors"
	"fmt"
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
	if !ok {
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
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
