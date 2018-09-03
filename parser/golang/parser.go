package golang

import (
	"go/token"
	"go/parser"
	"go/ast"
	"fmt"
	"strings"
	"github.com/jetuuuu/linter/diapason"
)

const (
	optional  = "_optional"
	linterTag = "js:"
)

type inspector struct {
	declFunctions map[string]string
	functions map[string]diapason.Range
}

func Parse(file string) (map[string]diapason.Range, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	i := &inspector{make(map[string]string), make(map[string]diapason.Range)}
	ast.Inspect(f, i.inspect)

	return i.functions, nil
}

func (i *inspector) inspect(n ast.Node) bool {
	call, ok := n.(*ast.CallExpr)
	if ok && call != nil && call.Fun != nil {
		var functionName string
		if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
			functionName = fun.Sel.Name
		} else {
			functionName = fmt.Sprintf("%s", call.Fun)
		}

		if comment, ok := i.declFunctions[functionName]; ok {
			args := diapason.Range{Min: len(call.Args), Max: len(call.Args)}
			if comment != "" {
				s := strings.TrimLeft(comment, linterTag)
				s = strings.TrimSpace(s)
				s = strings.Replace(s, "?", optional, -1)

				expr, err := parser.ParseExpr(s)
				if err == nil {
					if call, ok := expr.(*ast.CallExpr); ok {
						for _, a := range call.Args {
							if ident, ok := a.(*ast.Ident); ok {
								if strings.HasSuffix(ident.Name, optional) {
									args.Min--
								}
							}
						}
					}
				}
			}

			i.functions[functionName] = args
		}
	} else if decl, ok := n.(*ast.FuncDecl); ok {
		name := decl.Name.Name
		if decl.Name.IsExported() {
			var comment string
			if decl.Doc != nil {
				comment = decl.Doc.Text()
			}

			i.declFunctions[name] = comment
		}
	}

	return true
}