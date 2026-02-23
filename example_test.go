package url_test

import (
	"fmt"
	"log"

	"github.com/praffq/go-url"
)

func ExampleParse() {

	uri, err := url.Parse("https://example.com/sample_path")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uri)
	// Output: https://example.com/sample_path
}
