package simple

type invoker struct {}

func (i *invoker) Sum(a, b, c int) int {
	return a + b + c
}

var invoke = new(invoker)

//js: Sum(a, b, c?)
func Sum(a, b, c int) int {
	return 	invoke.Sum(a, b, c)
}
