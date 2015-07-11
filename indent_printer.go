package main

import (
	"fmt"
	"io"
)

type indentPrinter struct {
	Output       io.Writer
	IndentString string

	indentLevel int
	indented    bool
	err         error
}

func (p *indentPrinter) Indent() int {
	p.indentLevel++
	return p.indentLevel
}

func (p *indentPrinter) Outdent() int {
	if p.indentLevel > 0 {
		p.indentLevel--
	}
	return p.indentLevel
}

func (p *indentPrinter) printIndent() (n int, err error) {
	for i := 0; i < p.indentLevel; i++ {
		m, err := fmt.Print(p.IndentString)
		n += m
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (p *indentPrinter) Print(a ...interface{}) (n int, err error) {
	if p.err != nil {
		return 0, p.err
	}
	if !p.indented {
		p.indented = true
		m, err := p.printIndent()
		n += m
		if err != nil {
			return n, err
		}
	}
	n, p.err = fmt.Fprint(p.Output, a...)
	return n, p.err
}

func (p *indentPrinter) Printf(f string, a ...interface{}) (n int, err error) {
	if p.err != nil {
		return 0, p.err
	}
	if !p.indented {
		p.indented = true
		m, err := p.printIndent()
		n += m
		if err != nil {
			return n, err
		}
	}
	n, p.err = fmt.Fprintf(p.Output, f, a...)
	return n, p.err
}

func (p *indentPrinter) Println(a ...interface{}) (n int, err error) {
	if p.err != nil {
		return 0, p.err
	}
	if !p.indented {
		p.indented = true
		m, err := p.printIndent()
		n += m
		if err != nil {
			return n, err
		}
	}
	n, p.err = fmt.Fprintln(p.Output, a...)
	p.indented = false
	return n, p.err
}

func (p *indentPrinter) Error() error {
	return p.err
}
