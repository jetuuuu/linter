package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	otto_ast "github.com/robertkrimen/otto/ast"
	otto_Parser "github.com/robertkrimen/otto/parser"

	"flag"
	"time"
)

const (
	mayBeAbsent = "_mayBeAbsent"
	linterTag = "js:"
)

func main() {

	gojaApi := flag.String("go", "", "path to goja api .go file")
	js := flag.String("js", "", "path to js dir")
	flag.Parse()

	if err := validFlag(gojaApi, "go"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := validFlag(js, "js"); err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()

	m, err := parseGojaApi(*gojaApi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	problems := checkJS(*js, m)

	fmt.Printf("Time: %s\n", time.Since(start))
	fmt.Println(len(problems))
	fmt.Println(strings.Join(problems, "\n"))
}

func validFlag(f *string, name string) error {
	if f == nil || *f == "" {
		return fmt.Errorf("flag %s must be not empty", name)
	}

	return nil
}

func parseGojaApi(file string) (map[string]Range, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	m := make(map[string]Range)
	functions := make(map[string]string)


	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if ok && call != nil && call.Fun != nil {
			var functionName string
			if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
				functionName = fun.Sel.Name
			} else {
				functionName = fmt.Sprintf("%s", call.Fun)
			}

			if comment, ok := functions[functionName]; ok {
				args := Range{min: len(call.Args), max:len(call.Args)}
				if comment != "" {
					s := comment[strings.Index(comment, linterTag):]
					s = strings.Replace(s, "?", mayBeAbsent, -1)
					expr, err := parser.ParseExpr(s)
					if err == nil {
						if call, ok := expr.(*ast.CallExpr); ok {
							min := 0
							for _, a := range call.Args {
								if !strings.HasSuffix(a.(*ast.Ident).Name, mayBeAbsent) {
									min++
								}
							}
							args.min = min
						}
					} else {
						fmt.Println(err)
					}
				}

				m[functionName] = args
			}
		} else if decl, ok := n.(*ast.FuncDecl); ok {
			name := decl.Name.Name
			if decl.Name.IsExported() {
				var comment string
				if decl.Doc != nil {
					comment = decl.Doc.Text()
				}

				functions[name] = comment
			}
		}

		return true
	})

	return m, nil
}

func checkJS(dir string, m map[string]Range) []string {
	files := getAllFiles(dir)
	fmt.Println(len(files))

	var problems []string
	for i := range files {

		prg, err := otto_Parser.ParseFile(nil, files[i], nil, 0)
		if err != nil {
			continue
		}

		Inspect(prg, func(node otto_ast.Node) bool {
			call, ok := node.(*otto_ast.CallExpression)
			if ok {
				if ident, ok := call.Callee.(*otto_ast.Identifier); ok {
					actualArgs := len(call.ArgumentList)
					if expectedArgs, ok := m[ident.Name]; ok && !expectedArgs.Contains(actualArgs) {
						problems = append(problems, fmt.Sprintf("%s: %s must %s but %d", files[i], ident.Name, expectedArgs, actualArgs))
					}
				}
			}

			return true
		})
	}

	return problems
}
