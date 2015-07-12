package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

// GoPrinter is printer for golang.
var GoPrinter = &goPrinter{}

type goPrinter struct {
}

func (p *goPrinter) Fprint(output io.Writer, file *ast.File) error {
	fs := token.NewFileSet()
	return printer.Fprint(output, fs, file)
}
