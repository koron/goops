package main

import (
	"go/printer"
	"go/token"
	"io"
)

// GoPrinter is printer for golang.
var GoPrinter = &goPrinter{}

type goPrinter struct {
}

func (p *goPrinter) Fprint(output io.Writer, node interface{}) error {
	fs := token.NewFileSet()
	return printer.Fprint(output, fs, node)
}
