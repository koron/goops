package main

import (
	"go/ast"
	"io"
)

// JavaPrinter is printer for Java.
var JavaPrinter = &javaPrinter{}

type javaPrinter struct {
}

func (p *javaPrinter) Fprint(output io.Writer, list []ast.Stmt) error {
	// TODO:
	return nil
}
