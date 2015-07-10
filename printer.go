package main

import (
	"io"
	"strings"
)

type codePrinter interface {
	Fprint(output io.Writer, node interface{}) error
}

var printerMap = map[string]codePrinter{
	"go":   GoPrinter,
	"java": JavaPrinter,
}

func getPrinter(lang string) codePrinter {
	p, ok := printerMap[strings.ToLower(lang)]
	if !ok {
		return nil
	}
	return p
}
