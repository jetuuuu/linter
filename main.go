package main

import (
	"fmt"
	"flag"
	"github.com/jetuuuu/linter/parser/golang"
	"github.com/jetuuuu/linter/parser/js"
)

func main() {

	gojaApi := flag.String("go", "", "path to goja api .go file")
	jsDir := flag.String("js", "", "path to js dir")
	flag.Parse()

	if err := validFlag(gojaApi, "go"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := validFlag(jsDir, "js"); err != nil {
		fmt.Println(err.Error())
		return
	}

	gojaFunctions, err := golang.Parse(*gojaApi)
	if err != nil {
		panic(err.Error())
	}

	problems := js.ParseAndCheck(*jsDir, gojaFunctions)

	fmt.Println("Total errors:", len(problems))
	for _, p := range problems {
		fmt.Printf("%s:%d:%d -- %s must %s but %d\n", p.File, p.Pos.Line, p.Pos.Pos, p.Function, p.ExpectedArgs, p.ActualArgs)
	}
}

func validFlag(f *string, name string) error {
	if f == nil || *f == "" {
		return fmt.Errorf("flag %s must be not empty", name)
	}

	return nil
}
