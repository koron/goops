package main

import "io"

// JavaPrinter is printer for Java.
var JavaPrinter = &javaPrinter{}

type javaPrinter struct {
}

func (p *javaPrinter) Fprint(output io.Writer, node interface{}) error {
	// TODO:
	return nil
}
