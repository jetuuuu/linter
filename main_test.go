package main

import (
	"testing"
)

func Test_parseGojaApi(t *testing.T) {
	functions, err := parseGojaApi("testdata/external/go/simple.go")

	if err != nil {
		t.Fatalf("err must be nil but %s", err.Error())
	}

	if len(functions) == 0 {
		t.Fatal("result must not be empty")
	}

	f, ok := functions["Sum"]
	if !ok {
		t.Fatal("result must contain functin \"Sum\"")
	}

	r := Range{min: 2, max: 3}
	if f != r {
		t.Fatalf("must be %s but %s", r, f)
	}
}
