package js

import (
	"github.com/jetuuuu/linter/diapason"
	"io/ioutil"
	"path"
	"strings"

	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/file"
)

var cache = make(map[string][]string)

type Problem struct {
	File         string
	Pos          Position
	Function     string
	ActualArgs   int
	ExpectedArgs diapason.Range
}

type Position struct {
	Line int
	Pos int
}

func newPosition(lines []string, idx file.Idx) Position {
	p := Position{Line: 0, Pos: int(idx)}

	seek := 0
	for i, line := range lines {
		size := len(line)
		seek += size + 1
		if seek  >= p.Pos {
			p.Line = i + 1
			break
		}
	}

	return p
}

func ParseAndCheck(dir string, golangFunctions map[string]diapason.Range) []Problem {
	files := getAllFiles(dir)

	var problems []Problem
	for i := range files {
		prg, err := parser.ParseFile(nil, files[i], nil, 0)
		if err != nil {
			continue
		}

		lines := strings.Split(prg.File.Source(), "\n")

		Inspect(prg, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpression)
			if ok {
				if ident, ok := call.Callee.(*ast.Identifier); ok {
					actualArgs := len(call.ArgumentList)
					if expectedArgs, ok := golangFunctions[ident.Name]; ok && !expectedArgs.Contains(actualArgs) {
						problems = append(problems, Problem{
							File:         files[i],
							Pos:          newPosition(lines, ident.Idx0()),
							Function:     ident.Name,
							ActualArgs:   actualArgs,
							ExpectedArgs: expectedArgs,
						})
					}
				}
			}
			return true
		})
	}

	return problems
}


func getAllFiles(root string) []string {
	var ret []string
	list, _ := ioutil.ReadDir(root)
	for _, l := range list {
		if l.IsDir() {
			ret = append(ret, getAllFiles(path.Join(root, l.Name()))...)
		} else if strings.HasSuffix(l.Name(), ".js") {
			ret = append(ret, path.Join(root, l.Name()))
		}
	}

	return ret
}
