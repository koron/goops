package main

import (
	"go/ast"
	"io"
)

// DumpPrinter is printer for debug.
var DumpPrinter = &dumpPrinter{}

type dumpPrinter struct {
}

func (p *dumpPrinter) Fprint(w io.Writer, list []ast.Stmt) error {
	c := &dumpContext{
		codePrinterContext{
			output:       w,
			indentString: "  ",
		},
	}
	for _, s := range list {
		c.printStmt(s)
		if c.err != nil {
			break
		}
	}
	return c.err
}

type dumpContext struct {
	codePrinterContext
}

func (c *dumpContext) printBlockStmt(s *ast.BlockStmt) {
	for _, s := range s.List {
		c.printStmt(s)
	}
}

func (c *dumpContext) printIfStmt(s *ast.IfStmt) {
	c.Emit("IfStmt:")
	if s.Init != nil {
		c.Emit(" +Init:")
		c.Indent()
		c.printStmt(s.Init)
		c.Outdent()
	}
	if s.Cond != nil {
		c.Emit(" +Cond:")
		c.Indent()
		c.printExpr(s.Cond)
		c.Outdent()
	}
	if s.Body != nil {
		c.Emit(" +Body:")
		c.Indent()
		c.printBlockStmt(s.Body)
		c.Outdent()
	}
	if s.Else != nil {
		c.Emit(" +Else:")
		c.Indent()
		c.printStmt(s.Else)
		c.Outdent()
	}
}

func (c *dumpContext) printReturnStmt(s *ast.ReturnStmt) {
	c.Emit("ReturnStmt:")
	if s.Results != nil {
		c.Emit(" +Results:")
		c.Indent()
		for _, ex := range s.Results {
			c.printExpr(ex)
		}
		c.Outdent()
	}
}

func (c *dumpContext) printExpr(expr ast.Expr) {
	// TODO:
	c.Emit("(printExpr) %#v", expr)
}

func (c *dumpContext) printStmt(s ast.Stmt) {
	switch t := s.(type) {
	case *ast.IfStmt:
		c.printIfStmt(t)
	case *ast.ReturnStmt:
		c.printReturnStmt(t)
	// TODO: more Stmt
	default:
		c.Emit("(printStmt) %#v", s)
	}
}
