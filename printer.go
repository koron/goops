package main

import (
	"fmt"
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
	Fprint(output io.Writer, list []ast.Stmt) error
}

func getPrinter(lang string) codePrinter {
	p, ok := printerMap[strings.ToLower(lang)]
	if !ok {
		return nil
	}
	return p
}

type codePrinterContext struct {
	output       io.Writer
	indentLevel  int
	indentString string
	err          error
}

func (c *codePrinterContext) Indent() {
	c.indentLevel++
}

func (c *codePrinterContext) Outdent() {
	if c.indentLevel > 0 {
		c.indentLevel--
	}
}

func (c *codePrinterContext) Print(args ...interface{}) {
	if c.err != nil {
		return
	}
	_, c.err = fmt.Fprint(c.output, args...)
}

// Emit print a line with indent.
func (c *codePrinterContext) Emit(s string, args ...interface{}) {
	if c.err != nil {
		return
	}
	for i := 0; i < c.indentLevel; i++ {
		c.Print(c.indentString)
	}
	_, c.err = fmt.Fprintf(c.output, s, args...)
	c.Print("\n")
}
