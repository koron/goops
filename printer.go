package main

import (
	"go/ast"
	"io"
	"strings"
)

var printerMap = map[string]codePrinter{
	"go":   GoPrinter,
	"dump": DumpPrinter,
	"java": JavaPrinter,
}

type codePrinter interface {
	Fprint(output io.Writer, file *ast.File) error
}

func getPrinter(lang string) codePrinter {
	p, ok := printerMap[strings.ToLower(lang)]
	if !ok {
		return nil
	}
	return p
}
