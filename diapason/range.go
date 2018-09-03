package diapason

import "fmt"

type Range struct {
	Min int
	Max int
}

func (r Range) Contains(i int) bool {
	return i >= r.Min && i <= r.Max
}

func (r Range) String() string {
	if r.Min == r.Max {
		return fmt.Sprintf("%d", r.Max)
	} else {
		return fmt.Sprintf("between %d and %d", r.Min, r.Max)
	}
}
