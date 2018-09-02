package main

import (
	"io/ioutil"
	"path"
	"strings"

	otto_ast "github.com/robertkrimen/otto/ast"
)

type inspector func(otto_ast.Node) bool

func (f inspector) Enter(node otto_ast.Node) otto_ast.Visitor {
	if f(node) {
		return f
	}
	return nil
}

func (f inspector) Exit(n otto_ast.Node) {}

func Inspect(node otto_ast.Node, f func(otto_ast.Node) bool) {
	otto_ast.Walk(inspector(f), node)
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
