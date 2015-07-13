package main

import (
	"go/ast"
	"io"
	"strings"
)

var printerMap = map[string]codePrinter{}

type codePrinter interface {
	Fprint(output io.Writer, list []ast.Stmt) error
}

func getPrinter(lang string) codePrinter {
	p, ok := printerMap[strings.ToLower(lang)]
	if !ok {
		return nil
	}
	return p
}
