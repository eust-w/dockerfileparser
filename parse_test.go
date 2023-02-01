package dockerfileparser

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	d, err := ParseFile("./test/dockerfile")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	fmt.Println(d.All, len(d.All), len(d.All[0]))
}
