package main

import "fmt"

type Range struct {
	min int
	max int
}

func (r Range) Contains(i int) bool {
	return i >= r.min && i <= r.max
}

func (r Range) String() string {
	if r.min == r.max {
		return fmt.Sprintf("%d", r.max)
	} else {
		return fmt.Sprintf("between %d and %d", r.min, r.max)
	}
}
