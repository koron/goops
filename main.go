package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"io"
	"io/ioutil"
	"os"
)

const progname = "goops"

var (
	lang string
)

func main() {
	flag.StringVar(&lang, "l", "go", "output language")
	flag.Parse()
	err := goops(os.Stdin, os.Stdout, lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", progname, err)
		os.Exit(1)
	}
}

func strip(expr ast.Expr) ([]ast.Stmt, error) {
	f, ok := expr.(*ast.FuncLit)
	if !ok {
		return nil, errors.New("missing wrapper func")
	}
	return f.Body.List, nil
}

func goops(r io.Reader, w io.Writer, lang string) error {
	p := getPrinter(lang)
	if p == nil {
		return fmt.Errorf("unknown language: %s", lang)
	}
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read error: %s", err)
	}
	expr, err := parser.ParseExpr("func(){" + string(in) + "}")
	if err != nil {
		return fmt.Errorf("parse error: %s\n", err)
	}
	list, err := strip(expr)
	if err != nil {
		return fmt.Errorf("strip error: %s\n", err)
	}
	err = p.Fprint(w, list)
	if err != nil {
		return fmt.Errorf("write error: %s", err)
	}
	return nil
}
