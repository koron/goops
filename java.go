package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

// JavaPrinter is printer for Java.
var JavaPrinter = &javaPrinter{}

type javaPrinter struct {
}

func (p *javaPrinter) Fprint(output io.Writer, list []ast.Stmt) error {
	c := &javaContext{
		p: indentPrinter{
			Output:       output,
			IndentString: "    ",
		},
	}
	for _, x := range list {
		if err := c.printStmt(x); err != nil {
			return err
		}
	}
	return nil
}

type javaContext struct {
	p indentPrinter
	e error

	disableEOL bool
}

func (c *javaContext) err() error {
	if c.e != nil {
		return c.e
	}
	return c.p.Error()
}

// dummy function
func (c *javaContext) nodeStart(n string) {
	c.p.Printf("[[%s]]", n)
}

// dummy function
func (c *javaContext) nodeEnd() error {
	if c.e != nil {
		return c.e
	}
	return c.p.Error()
}

// dummy function
func (c *javaContext) emitProp(n string, v interface{}) {
}

////////////////////////////////////////////////////////////////////////////
// Expr

func (c *javaContext) printExpr(expr ast.Expr) (err error) {
	switch x := expr.(type) {
	case *ast.BadExpr:
		err = c.printBadExpr(x)
	case *ast.Ident:
		err = c.printIdent(x)
	case *ast.Ellipsis:
		err = c.printEllipsis(x)
	case *ast.BasicLit:
		err = c.printBasicLit(x)
	case *ast.FuncLit:
		err = c.printFuncLit(x)
	case *ast.CompositeLit:
		err = c.printCompositeLit(x)
	case *ast.ParenExpr:
		err = c.printParenExpr(x)
	case *ast.SelectorExpr:
		err = c.printSelectorExpr(x)
	case *ast.IndexExpr:
		err = c.printIndexExpr(x)
	case *ast.SliceExpr:
		err = c.printSliceExpr(x)
	case *ast.TypeAssertExpr:
		err = c.printTypeAssertExpr(x)
	case *ast.CallExpr:
		err = c.printCallExpr(x)
	case *ast.StarExpr:
		err = c.printStarExpr(x)
	case *ast.UnaryExpr:
		err = c.printUnaryExpr(x)
	case *ast.BinaryExpr:
		err = c.printBinaryExpr(x)
	case *ast.KeyValueExpr:
		err = c.printKeyValueExpr(x)
	case *ast.ArrayType:
		err = c.printArrayType(x)
	case *ast.StructType:
		err = c.printStructType(x)
	case *ast.FuncType:
		err = c.printFuncType(x)
	case *ast.InterfaceType:
		err = c.printInterfaceType(x)
	case *ast.MapType:
		err = c.printMapType(x)
	case *ast.ChanType:
		err = c.printChanType(x)
	default:
		return fmt.Errorf("unsupported expr: %#v", expr)
	}
	return err
}

func (c *javaContext) printBadExpr(x *ast.BadExpr) error {
	c.nodeStart("BadExpr")
	return c.nodeEnd()
}

func (c *javaContext) printIdent(x *ast.Ident) error {
	c.p.Print(x.Name)
	// TODO: consider x.Obj
	return c.err()
}

func (c *javaContext) printEllipsis(x *ast.Ellipsis) error {
	c.nodeStart("Ellipsis")
	c.emitProp("Elt", x.Elt)
	return c.nodeEnd()
}

func (c *javaContext) printBasicLit(x *ast.BasicLit) error {
	c.p.Print(x.Value)
	// TODO: consider x.Kind
	return c.err()
}

func (c *javaContext) printFuncLit(x *ast.FuncLit) error {
	c.nodeStart("FuncLit")
	c.emitProp("Type", x.Type)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printCompositeLit(x *ast.CompositeLit) error {
	c.nodeStart("CompositeLit")
	c.emitProp("Type", x.Type)
	c.emitProp("Elts", x.Elts)
	return c.nodeEnd()
}

func (c *javaContext) printParenExpr(x *ast.ParenExpr) error {
	c.nodeStart("ParenExpr")
	c.emitProp("X", x.X)
	return c.nodeEnd()
}

func (c *javaContext) printSelectorExpr(x *ast.SelectorExpr) error {
	c.nodeStart("SelectorExpr")
	c.emitProp("X", x.X)
	c.emitProp("Sel", x.Sel)
	return c.nodeEnd()
}

func (c *javaContext) printIndexExpr(x *ast.IndexExpr) error {
	c.nodeStart("IndexExpr")
	c.emitProp("X", x.X)
	c.emitProp("Index", x.Index)
	return c.nodeEnd()
}

func (c *javaContext) printSliceExpr(x *ast.SliceExpr) error {
	c.nodeStart("SliceExpr")
	c.emitProp("X", x.X)
	c.emitProp("Low", x.Low)
	c.emitProp("High", x.High)
	c.emitProp("Max", x.Max)
	c.emitProp("Slice3", x.Slice3)
	return c.nodeEnd()
}

func (c *javaContext) printTypeAssertExpr(x *ast.TypeAssertExpr) error {
	c.nodeStart("TypeAssertExpr")
	c.emitProp("X", x.X)
	c.emitProp("Type", x.Type)
	return c.nodeEnd()
}

func (c *javaContext) printCallExpr(x *ast.CallExpr) error {
	c.nodeStart("CallExpr")
	c.emitProp("Func", x.Fun)
	c.emitProp("Args", x.Args)
	return c.nodeEnd()
}

func (c *javaContext) printStarExpr(x *ast.StarExpr) error {
	c.nodeStart("StarExpr")
	c.emitProp("X", x.X)
	return c.nodeEnd()
}

func (c *javaContext) printUnaryExpr(x *ast.UnaryExpr) error {
	c.p.Print(x.Op)
	c.printExpr(x.X)
	return c.err()
}

func (c *javaContext) printBinaryExpr(x *ast.BinaryExpr) error {
	// special case: String#equals()
	if y, ok := x.Y.(*ast.BasicLit); (x.Op == token.EQL || x.Op == token.NEQ) &&
		ok && y.Kind == token.STRING {
		if x.Op == token.NEQ {
			c.p.Print("!")
		}
		c.printExpr(x.X)
		c.p.Print(".equals(")
		c.printExpr(x.Y)
		c.p.Print(")")
		return c.err()
	}
	// normal case
	c.printExpr(x.X)
	c.p.Printf(" %s ", x.Op)
	c.printExpr(x.Y)
	return c.err()
}

func (c *javaContext) printKeyValueExpr(x *ast.KeyValueExpr) error {
	c.nodeStart("KeyValueExpr")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	return c.nodeEnd()
}

func (c *javaContext) printArrayType(x *ast.ArrayType) error {
	c.nodeStart("ArrayType")
	c.emitProp("Len", x.Len)
	c.emitProp("Elt", x.Elt)
	return c.nodeEnd()
}

func (c *javaContext) printStructType(x *ast.StructType) error {
	c.nodeStart("StructType")
	c.emitProp("FieldList", x.Fields)
	c.emitProp("Incomplete", x.Incomplete)
	return c.nodeEnd()
}

func (c *javaContext) printFuncType(x *ast.FuncType) error {
	c.nodeStart("FuncType")
	c.emitProp("Params", x.Params)
	c.emitProp("Results", x.Results)
	return c.nodeEnd()
}

func (c *javaContext) printInterfaceType(x *ast.InterfaceType) error {
	c.nodeStart("InterfaceType")
	c.emitProp("Methods", x.Methods)
	c.emitProp("Incomplete", x.Incomplete)
	return c.nodeEnd()
}

func (c *javaContext) printMapType(x *ast.MapType) error {
	c.nodeStart("MapType")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	return c.nodeEnd()
}

func (c *javaContext) printChanType(x *ast.ChanType) error {
	c.nodeStart("ChanType")
	c.emitProp("Dir", x.Dir)
	c.emitProp("Value", x.Value)
	return c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Stmt

func (c *javaContext) printStmt(stmt ast.Stmt) (err error) {
	switch x := stmt.(type) {
	case *ast.BadStmt:
		err = c.printBadStmt(x)
	case *ast.DeclStmt:
		err = c.printDeclStmt(x)
	case *ast.EmptyStmt:
		err = c.printEmptyStmt(x)
	case *ast.LabeledStmt:
		err = c.printLabeledStmt(x)
	case *ast.ExprStmt:
		err = c.printExprStmt(x)
	case *ast.SendStmt:
		err = c.printSendStmt(x)
	case *ast.IncDecStmt:
		err = c.printIncDecStmt(x)
	case *ast.AssignStmt:
		err = c.printAssignStmt(x)
	case *ast.GoStmt:
		err = c.printGoStmt(x)
	case *ast.DeferStmt:
		err = c.printDeferStmt(x)
	case *ast.ReturnStmt:
		err = c.printReturnStmt(x)
	case *ast.BranchStmt:
		err = c.printBranchStmt(x)
	case *ast.BlockStmt:
		err = c.printBlockStmt(x)
	case *ast.IfStmt:
		err = c.printIfStmt(x)
	case *ast.CaseClause:
		err = c.printCaseClause(x)
	case *ast.SwitchStmt:
		err = c.printSwitchStmt(x)
	case *ast.TypeSwitchStmt:
		err = c.printTypeSwitchStmt(x)
	case *ast.CommClause:
		err = c.printCommClause(x)
	case *ast.SelectStmt:
		err = c.printSelectStmt(x)
	case *ast.ForStmt:
		err = c.printForStmt(x)
	case *ast.RangeStmt:
		err = c.printRangeStmt(x)
	default:
		return fmt.Errorf("unsupported Stmt: %#v", stmt)
	}
	return err
}

func (c *javaContext) printBadStmt(x *ast.BadStmt) error {
	c.nodeStart("BadStmt")
	return c.nodeEnd()
}

func (c *javaContext) printDeclStmt(x *ast.DeclStmt) error {
	c.nodeStart("DeclStmt")
	c.emitProp("Decl", x.Decl)
	return c.nodeEnd()
}

func (c *javaContext) printEmptyStmt(x *ast.EmptyStmt) error {
	c.nodeStart("EmptyStmt")
	return c.nodeEnd()
}

func (c *javaContext) printLabeledStmt(x *ast.LabeledStmt) error {
	c.nodeStart("LabeledStmt")
	c.emitProp("Label", x.Label)
	c.emitProp("Stmt", x.Stmt)
	return c.nodeEnd()
}

func (c *javaContext) printExprStmt(x *ast.ExprStmt) error {
	c.printExpr(x.X)
	c.p.Println()
	return c.err()
}

func (c *javaContext) printSendStmt(x *ast.SendStmt) error {
	c.nodeStart("SendStmt")
	c.emitProp("Chan", x.Chan)
	c.emitProp("Value", x.Value)
	return c.nodeEnd()
}

func (c *javaContext) printIncDecStmt(x *ast.IncDecStmt) error {
	c.printExpr(x.X)
	c.p.Print(x.Tok)
	c.printEOL()
	return c.err()
}

func (c *javaContext) printAssignStmt(x *ast.AssignStmt) error {
	for i, lhs := range x.Lhs {
		c.printAssign(lhs, x.Rhs[i], x.Tok)
		c.printEOL()
	}
	return c.nodeEnd()
}

func (c *javaContext) printGoStmt(x *ast.GoStmt) error {
	c.nodeStart("GoStmt")
	c.emitProp("Call", x.Call)
	return c.nodeEnd()
}

func (c *javaContext) printDeferStmt(x *ast.DeferStmt) error {
	c.nodeStart("DeferStmt")
	c.emitProp("Call", x.Call)
	return c.nodeEnd()
}

func (c *javaContext) printReturnStmt(x *ast.ReturnStmt) error {
	switch len(x.Results) {
	case 0:
		c.p.Print("return")
	case 1:
		c.p.Print("return ")
		c.printExpr(x.Results[0])
	default:
		c.p.Println("// FIXME: return accepts only one value.")
		c.p.Print("return (")
		first := true
		for _, expr := range x.Results {
			if first {
				first = false
			} else {
				c.p.Print(", ")
			}
			c.printExpr(expr)
		}
		c.p.Print(")")
	}
	c.printEOL()
	return c.err()
}

func (c *javaContext) printBranchStmt(x *ast.BranchStmt) error {
	c.nodeStart("BranchStmt")
	c.emitProp("Tok", x.Tok)
	c.emitProp("Label", x.Label)
	return c.nodeEnd()
}

func (c *javaContext) printBlockStmt(x *ast.BlockStmt) error {
	for _, stmt := range x.List {
		if err := c.printStmt(stmt); err != nil {
			return err
		}
	}
	return c.err()
}

func (c *javaContext) printIfStmt(x *ast.IfStmt) error {
	if x.Init != nil {
		c.printStmt(x.Init)
	}
	c.p.Print("if (")
	if x.Cond != nil {
		c.printExpr(x.Cond)
	}
	c.p.Println(") {")
	if x.Body != nil {
		c.p.Indent()
		c.printStmt(x.Body)
		c.p.Outdent()
	}
	if x.Else != nil {
		c.p.Println("} else {")
		c.p.Indent()
		c.printStmt(x.Body)
		c.p.Outdent()
	}
	c.p.Println("}")
	return c.err()
}

func (c *javaContext) printCaseClause(x *ast.CaseClause) error {
	c.nodeStart("CaseClause")
	c.emitProp("List", x.List)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printSwitchStmt(x *ast.SwitchStmt) error {
	c.nodeStart("SwitchStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Tag", x.Tag)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printTypeSwitchStmt(x *ast.TypeSwitchStmt) error {
	c.nodeStart("TypeSwitchStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Assign", x.Assign)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printCommClause(x *ast.CommClause) error {
	c.nodeStart("CommClause")
	c.emitProp("Comm", x.Comm)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printSelectStmt(x *ast.SelectStmt) error {
	c.nodeStart("SelectStmt")
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

func (c *javaContext) printForStmt(x *ast.ForStmt) error {
	if x.Init == nil && x.Post == nil {
		c.p.Print("while (")
		if x.Cond != nil {
			c.printExpr(x.Cond)
		} else {
			c.p.Print("true")
		}
		c.p.Println(") {")
	} else {
		c.p.Print("for (")
		c.disableEOL = true
		if x.Init != nil {
			c.printStmt(x.Init)
		}
		c.p.Print(";")
		if x.Cond != nil {
			c.p.Print(" ")
			c.printExpr(x.Cond)
		}
		c.p.Print(";")
		if x.Post != nil {
			c.p.Print(" ")
			c.printStmt(x.Post)
		}
		c.disableEOL = false
		c.p.Println(") {")
	}
	c.p.Indent()
	c.printBlockStmt(x.Body)
	c.p.Outdent()
	c.p.Println("}")
	return c.err()
}

func (c *javaContext) printRangeStmt(x *ast.RangeStmt) error {
	c.nodeStart("RangeStmt")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	c.emitProp("Tok", x.Tok)
	c.emitProp("X", x.X)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Decl

func (c *javaContext) printDecl(decl ast.Decl) (err error) {
	switch x := decl.(type) {
	case *ast.BadDecl:
		err = c.printBadDecl(x)
	case *ast.GenDecl:
		err = c.printGenDecl(x)
	case *ast.FuncDecl:
		err = c.printFuncDecl(x)
	default:
		return fmt.Errorf("unsupported decl: %#v", decl)
	}
	return err
}

func (c *javaContext) printBadDecl(x *ast.BadDecl) error {
	c.nodeStart("BadDecl")
	return c.nodeEnd()
}

func (c *javaContext) printGenDecl(x *ast.GenDecl) error {
	c.nodeStart("GenDecl")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Tok", x.Tok)
	c.emitProp("Specs", x.Specs)
	return c.nodeEnd()
}

func (c *javaContext) printFuncDecl(x *ast.FuncDecl) error {
	c.nodeStart("FuncDecl")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Recv", x.Recv)
	c.emitProp("Name", x.Name)
	c.emitProp("Type", x.Type)
	c.emitProp("Body", x.Body)
	return c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Spec

func (c *javaContext) printSpec(spec ast.Spec) (err error) {
	switch x := spec.(type) {
	case *ast.ImportSpec:
		err = c.printImportSpec(x)
	case *ast.ValueSpec:
		err = c.printValueSpec(x)
	case *ast.TypeSpec:
		err = c.printTypeSpec(x)
	default:
		return fmt.Errorf("unsupported spec: %#v", spec)
	}
	return err
}

func (c *javaContext) printImportSpec(x *ast.ImportSpec) error {
	c.nodeStart("ImportSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Path", x.Path)
	c.emitProp("Comment", x.Comment)
	return c.nodeEnd()
}

func (c *javaContext) printValueSpec(x *ast.ValueSpec) error {
	c.nodeStart("ValueSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Names", x.Names)
	c.emitProp("Type", x.Type)
	c.emitProp("Values", x.Values)
	c.emitProp("Comment", x.Comment)
	return c.nodeEnd()
}

func (c *javaContext) printTypeSpec(x *ast.TypeSpec) error {
	c.nodeStart("TypeSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Type", x.Type)
	c.emitProp("Comment", x.Comment)
	return c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Node

func (c *javaContext) printNode(node ast.Node) (err error) {
	switch x := node.(type) {
	case ast.Expr:
		err = c.printExpr(x)
	case ast.Stmt:
		err = c.printStmt(x)
	case ast.Decl:
		err = c.printDecl(x)
	case ast.Spec:
		err = c.printSpec(x)
	case *ast.Comment:
		err = c.printComment(x)
	case *ast.CommentGroup:
		err = c.printCommentGroup(x)
	case *ast.Field:
		err = c.printField(x)
	case *ast.FieldList:
		err = c.printFieldList(x)
	case *ast.File:
		err = c.printFile(x)
	case *ast.Package:
		err = c.printPackage(x)
	default:
		return fmt.Errorf("unsupported node: %#v", node)
	}
	return err
}

func (c *javaContext) printComment(x *ast.Comment) error {
	c.nodeStart("Comment")
	c.emitProp("Text", x.Text)
	return c.nodeEnd()
}

func (c *javaContext) printCommentGroup(x *ast.CommentGroup) error {
	c.nodeStart("CommentGroup")
	c.emitProp("List", x.List)
	return c.nodeEnd()
}

func (c *javaContext) printField(x *ast.Field) error {
	c.nodeStart("Field")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Names", x.Names)
	c.emitProp("Type", x.Type)
	c.emitProp("Tag", x.Tag)
	c.emitProp("Comment", x.Comment)
	return c.nodeEnd()
}

func (c *javaContext) printFieldList(x *ast.FieldList) error {
	c.nodeStart("FieldList")
	c.emitProp("List", x.List)
	return c.nodeEnd()
}

func (c *javaContext) printFile(x *ast.File) error {
	c.nodeStart("File")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Decls", x.Decls)
	c.emitProp("Scope", x.Scope)
	c.emitProp("Imports", x.Imports)
	c.emitProp("Unresolved", x.Unresolved)
	c.emitProp("Comments", x.Comments)
	return c.nodeEnd()
}

func (c *javaContext) printPackage(x *ast.Package) error {
	c.nodeStart("Package")
	c.emitProp("Name", x.Name)
	c.emitProp("Scope", x.Scope)
	c.emitProp("Imports", x.Imports)
	c.emitProp("Files", x.Files)
	return c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Others

func (c *javaContext) printEOL() {
	if c.disableEOL {
		return
	}
	c.p.Println(";")
}

func (c *javaContext) printType(exprs ...ast.Expr) {
	// TODO:
	c.p.Print("var")
}

func (c *javaContext) printAssign(lhs, rhs ast.Expr, tok token.Token) {
	switch tok {
	case token.DEFINE:
		c.printType(lhs, rhs)
		c.p.Print(" ")
		tok = token.ASSIGN
		fallthrough
	default:
		c.printExpr(lhs)
		c.p.Printf(" %s ", tok)
		c.printExpr(rhs)
	}
}
