package main

import (
	"testing"
)

func Test_parseGojaApi(t *testing.T) {
	functions, err := parseGojaApi("testdata/go/simple.go")

	if err != nil {
		t.Fatalf("err must be nil but %s", err.Error())
	}

	if len(functions) == 0 {
		t.Fatal("result must not be empty")
	}
}
