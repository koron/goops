package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

func init() {
	printerMap["go"] = GoPrinter
}

// GoPrinter is printer for golang.
var GoPrinter = &goPrinter{}

type goPrinter struct {
}

func (p *goPrinter) Fprint(output io.Writer, list []ast.Stmt) error {
	fs := token.NewFileSet()
	return printer.Fprint(output, fs, list)
}
