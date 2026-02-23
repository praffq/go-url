package url

import (
	"fmt"
	"strings"
	"testing"
)

func TestURLString(t *testing.T) {

	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{name: "nil", uri: nil, want: ""},
		{name: "empty", uri: &URL{}, want: ""},
		{name: "scheme_only", uri: &URL{Scheme: "http"}, want: "http:"},
		{name: "scheme_and_host", uri: &URL{Scheme: "http", Host: "example.com"}, want: "http://example.com"},
		{name: "scheme_host_and_path", uri: &URL{Scheme: "http", Host: "example.com", Path: "index.html"}, want: "http://example.com/index.html"},
		{name: "scheme_and_path", uri: &URL{Scheme: "data", Path: "text/plain"}, want: "data:text/plain"},
		{name: "host_only", uri: &URL{Host: "example.com"}, want: "example.com"},
		{name: "path_only", uri: &URL{Path: "index.html"}, want: "index.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			got := tt.uri.String()
			if got != tt.want {
				t.Errorf("URL.String() = %q, want %q", got, tt.want)
			}
		})
	}

}

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{name: "full", uri: "http://example.com/sample", want: &URL{Scheme: "http", Host: "example.com", Path: "sample"}},
	{name: "without_path", uri: "http://example.com/", want: &URL{Scheme: "http", Host: "example.com", Path: ""}},
	{name: "with_data_schema", uri: "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==", want: &URL{Scheme: "data"}},
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v, want <nil>", tt.uri, err)
			}
			if *got != *tt.want {
				t.Errorf("Parse(%q) = %v, want %v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{name: "missing_scheme", uri: "example.com/sample"},
		{name: "empty_schema", uri: ":example.com/sample"},
		{name: "with_data_schema", uri: "text/plain;base64,SGVsbG8sIFdvcmxkIQ=="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			_, err := Parse(tt.uri)
			if err == nil {
				t.Fatalf("Parse(%q) returned <nil>, want error 'missing scheme in URL'", tt.uri)
			}
		})
	}
}

func BenchmarkURLString(b *testing.B) {

	for _, n := range []int{10, 100, 1000, 10000} {
		u := &URL{Scheme: strings.Repeat("x", n), Host: strings.Repeat("y", n), Path: strings.Repeat("z", n)}

		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for b.Loop() {
				_ = u.String()
			}
		})

	}
}
