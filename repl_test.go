package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepl(t *testing.T) {
	test := map[string]struct {
		input    string
		expected []string
	}{
		"hw": {input: "  hello  woRld  ", expected: []string{"hello", "world"}},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf("%v\n", diff)
			}
		})
	}
}
