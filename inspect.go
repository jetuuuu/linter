package main

import (
	ast "github.com/robertkrimen/otto/ast"
)

type inspector func(ast.Node) bool

func (f inspector) Enter(node ast.Node) ast.Visitor {
	if f(node) {
		return f
	}
	return nil
}

func (f inspector) Exit(n ast.Node) {}

func Inspect(node ast.Node, f func(ast.Node) bool) {
	ast.Walk(inspector(f), node)
}