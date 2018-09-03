package testdata

type invoker struct {}

func (i invoker) Sum(a, b, c int) int {
	return a + b + c
}

func (i invoker) IncOne(one, value int) int {
	return one * value
}

var invoke = new(invoker)

//js: Sum(a, b, c?)
func Sum(a, b, c int) int {
	return 	invoke.Sum(a, b, c)
}


func IncOne(one, value int) {
	if value == 0 {
		value = 1
	}
	invoke.IncOne(one, value)
}