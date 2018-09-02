package main

import "testing"

func TestRange_Contains(t *testing.T) {
	r := Range{min: 100, max: 500}

	if !r.Contains(200) {
		t.Fatal("200 enters the interval between 100 and 500")
	}

	if r.Contains(0) {
		t.Fatal("0 doesn't enters the interval between 100 and 500")
	}
}

func TestRange_String(t *testing.T) {
	r := Range{100, 500}
	s := r.String()
	expected := "between 100 and 500"

	if s != expected {
		t.Fatalf("[min and max are not equal] must %s but %s", expected, s)
	}

	r1 := Range{min: 40, max: 40}
	s = r1.String()
	expected = "40"

	if s != expected {
		t.Fatalf("[min and max are equal] must %s but %s", expected, s)
	}
}
