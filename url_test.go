package url

import "testing"

func TestURLString(t *testing.T) {

	uri := &URL{Scheme: "http", Host: "example.com", Path: "sample"}

	got := uri.String()
	want := "http://example.com/sample"
	if got != want {
		t.Errorf("URL.String() = %q, want %q", got, want)
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
