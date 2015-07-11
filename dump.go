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
	for _, stmt := range list {
		c.printStmt(stmt)
		if c.err != nil {
			break
		}
	}
	return c.err
}

type dumpContext struct {
	codePrinterContext
}

func (c *dumpContext) nodeStart(n string) {
	c.Emit("%s {", n)
}

func (c *dumpContext) nodeEnd() {
	c.Emit("}")
}

func (c *dumpContext) propStart(n string) {
	c.Emit("+-%s:", n)
	c.Indent()
}

func (c *dumpContext) propEnd() {
	c.Outdent()
}

func (c *dumpContext) emitProp(n string, v interface{}) {
	if v == nil {
		return
	}
	if node, ok := v.(ast.Node); ok {
		c.propStart(n)
		c.printNode(node)
		c.propEnd()
		return
	}
	if array, ok := v.([]ast.Expr); ok {
		if len(array) > 0 {
			c.propStart(n)
			for _, item := range array {
				c.printExpr(item)
			}
			c.propEnd()
		}
		return
	}
	if array, ok := v.([]ast.Stmt); ok {
		if len(array) > 0 {
			c.propStart(n)
			for _, item := range array {
				c.printStmt(item)
			}
			c.propEnd()
		}
		return
	}
	c.Emit("+-%s: %v", n, v)
}

func (c *dumpContext) emitPropFn(n string, f func()) {
	c.propStart(n)
	f()
	c.propEnd()
}

////////////////////////////////////////////////////////////////////////////
// Expr

func (c *dumpContext) printExpr(expr ast.Expr) {
	switch e := expr.(type) {
	case *ast.BadExpr:
		c.printBadExpr(e)
	case *ast.Ident:
		c.printIdent(e)
	case *ast.Ellipsis:
		c.printEllipsis(e)
	case *ast.BasicLit:
		c.printBasicLit(e)
	case *ast.FuncLit:
		c.printFuncLit(e)
	case *ast.CompositeLit:
		c.printCompositeLit(e)
	case *ast.ParenExpr:
		c.printParenExpr(e)
	case *ast.SelectorExpr:
		c.printSelectorExpr(e)
	case *ast.IndexExpr:
		c.printIndexExpr(e)
	case *ast.SliceExpr:
		c.printSliceExpr(e)
	case *ast.TypeAssertExpr:
		c.printTypeAssertExpr(e)
	case *ast.CallExpr:
		c.printCallExpr(e)
	case *ast.StarExpr:
		c.printStarExpr(e)
	case *ast.UnaryExpr:
		c.printUnaryExpr(e)
	case *ast.BinaryExpr:
		c.printBinaryExpr(e)
	case *ast.KeyValueExpr:
		c.printKeyValueExpr(e)
	case *ast.ArrayType:
		c.printArrayType(e)
	case *ast.StructType:
		c.printStructType(e)
	case *ast.FuncType:
		c.printFuncType(e)
	case *ast.InterfaceType:
		c.printInterfaceType(e)
	case *ast.MapType:
		c.printMapType(e)
	case *ast.ChanType:
		c.printChanType(e)
	default:
		c.Emit("(Expr) %#v", expr)
	}
}

func (c *dumpContext) printBadExpr(expr *ast.BadExpr) {
	// TODO:
	c.Emit("(BadExpr) %#v", expr)
}

func (c *dumpContext) printIdent(expr *ast.Ident) {
	c.nodeStart("Ident")
	c.emitProp("Name", expr.Name)
	c.emitProp("Obj", expr.Obj)
	c.nodeEnd()
}

func (c *dumpContext) printEllipsis(expr *ast.Ellipsis) {
	// TODO:
	c.Emit("(Ellipsis) %#v", expr)
}

func (c *dumpContext) printBasicLit(expr *ast.BasicLit) {
	c.nodeStart("BasicLit")
	c.emitProp("Kind", expr.Kind)
	c.emitProp("Value", expr.Value)
	c.nodeEnd()
}

func (c *dumpContext) printFuncLit(expr *ast.FuncLit) {
	// TODO:
	c.Emit("(FuncLit) %#v", expr)
}

func (c *dumpContext) printCompositeLit(expr *ast.CompositeLit) {
	// TODO:
	c.Emit("(CompositeLit) %#v", expr)
}

func (c *dumpContext) printParenExpr(expr *ast.ParenExpr) {
	// TODO:
	c.Emit("(ParenExpr) %#v", expr)
}

func (c *dumpContext) printSelectorExpr(expr *ast.SelectorExpr) {
	// TODO:
	c.Emit("(SelectorExpr) %#v", expr)
}

func (c *dumpContext) printIndexExpr(expr *ast.IndexExpr) {
	// TODO:
	c.Emit("(IndexExpr) %#v", expr)
}

func (c *dumpContext) printSliceExpr(expr *ast.SliceExpr) {
	// TODO:
	c.Emit("(SliceExpr) %#v", expr)
}

func (c *dumpContext) printTypeAssertExpr(expr *ast.TypeAssertExpr) {
	// TODO:
	c.Emit("(TypeAssertExpr) %#v", expr)
}

func (c *dumpContext) printCallExpr(expr *ast.CallExpr) {
	// TODO:
	c.Emit("(CallExpr) %#v", expr)
}

func (c *dumpContext) printStarExpr(expr *ast.StarExpr) {
	// TODO:
	c.Emit("(StarExpr) %#v", expr)
}

func (c *dumpContext) printUnaryExpr(expr *ast.UnaryExpr) {
	// TODO:
	c.Emit("(UnaryExpr) %#v", expr)
}

func (c *dumpContext) printBinaryExpr(expr *ast.BinaryExpr) {
	c.nodeStart("BinaryExpr")
	c.emitProp("X", expr.X)
	c.emitProp("Op", expr.Op)
	c.emitProp("Y", expr.Y)
	c.nodeEnd()
}

func (c *dumpContext) printKeyValueExpr(expr *ast.KeyValueExpr) {
	// TODO:
	c.Emit("(KeyValueExpr) %#v", expr)
}

func (c *dumpContext) printArrayType(expr *ast.ArrayType) {
	// TODO:
	c.Emit("(ArrayType) %#v", expr)
}

func (c *dumpContext) printStructType(expr *ast.StructType) {
	// TODO:
	c.Emit("(StructType) %#v", expr)
}

func (c *dumpContext) printFuncType(expr *ast.FuncType) {
	// TODO:
	c.Emit("(FuncType) %#v", expr)
}

func (c *dumpContext) printInterfaceType(expr *ast.InterfaceType) {
	// TODO:
	c.Emit("(InterfaceType) %#v", expr)
}

func (c *dumpContext) printMapType(expr *ast.MapType) {
	// TODO:
	c.Emit("(MapType) %#v", expr)
}

func (c *dumpContext) printChanType(expr *ast.ChanType) {
	// TODO:
	c.Emit("(ChanType) %#v", expr)
}

////////////////////////////////////////////////////////////////////////////
// Stmt

func (c *dumpContext) printStmt(stmt ast.Stmt) {
	switch s := stmt.(type) {
	case *ast.BadStmt:
		c.printBadStmt(s)
	case *ast.DeclStmt:
		c.printDeclStmt(s)
	case *ast.EmptyStmt:
		c.printEmptyStmt(s)
	case *ast.LabeledStmt:
		c.printLabeledStmt(s)
	case *ast.ExprStmt:
		c.printExprStmt(s)
	case *ast.SendStmt:
		c.printSendStmt(s)
	case *ast.IncDecStmt:
		c.printIncDecStmt(s)
	case *ast.AssignStmt:
		c.printAssignStmt(s)
	case *ast.GoStmt:
		c.printGoStmt(s)
	case *ast.DeferStmt:
		c.printDeferStmt(s)
	case *ast.ReturnStmt:
		c.printReturnStmt(s)
	case *ast.BranchStmt:
		c.printBranchStmt(s)
	case *ast.BlockStmt:
		c.printBlockStmt(s)
	case *ast.IfStmt:
		c.printIfStmt(s)
	case *ast.CaseClause:
		c.printCaseClause(s)
	case *ast.SwitchStmt:
		c.printSwitchStmt(s)
	case *ast.TypeSwitchStmt:
		c.printTypeSwitchStmt(s)
	case *ast.CommClause:
		c.printCommClause(s)
	case *ast.SelectStmt:
		c.printSelectStmt(s)
	case *ast.ForStmt:
		c.printForStmt(s)
	case *ast.RangeStmt:
		c.printRangeStmt(s)
	default:
		c.Emit("(Stmt) %#v", stmt)
	}
}

func (c *dumpContext) printBadStmt(stmt *ast.BadStmt) {
	// TODO:
	c.Emit("(BadStmt) %#v", stmt)
}

func (c *dumpContext) printDeclStmt(stmt *ast.DeclStmt) {
	// TODO:
	c.Emit("(DeclStmt) %#v", stmt)
}

func (c *dumpContext) printEmptyStmt(stmt *ast.EmptyStmt) {
	// TODO:
	c.Emit("(EmptyStmt) %#v", stmt)
}

func (c *dumpContext) printLabeledStmt(stmt *ast.LabeledStmt) {
	// TODO:
	c.Emit("(LabeledStmt) %#v", stmt)
}

func (c *dumpContext) printExprStmt(stmt *ast.ExprStmt) {
	// TODO:
	c.Emit("(ExprStmt) %#v", stmt)
}

func (c *dumpContext) printSendStmt(stmt *ast.SendStmt) {
	// TODO:
	c.Emit("(SendStmt) %#v", stmt)
}

func (c *dumpContext) printIncDecStmt(stmt *ast.IncDecStmt) {
	// TODO:
	c.Emit("(IncDecStmt) %#v", stmt)
}

func (c *dumpContext) printAssignStmt(stmt *ast.AssignStmt) {
	// TODO:
	c.Emit("(AssignStmt) %#v", stmt)
}

func (c *dumpContext) printGoStmt(stmt *ast.GoStmt) {
	// TODO:
	c.Emit("(GoStmt) %#v", stmt)
}

func (c *dumpContext) printDeferStmt(stmt *ast.DeferStmt) {
	// TODO:
	c.Emit("(DeferStmt) %#v", stmt)
}

func (c *dumpContext) printReturnStmt(stmt *ast.ReturnStmt) {
	c.nodeStart("ReturnStmt")
	c.emitProp("Results", stmt.Results)
	c.nodeEnd()
}

func (c *dumpContext) printBranchStmt(stmt *ast.BranchStmt) {
	// TODO:
	c.Emit("(BranchStmt) %#v", stmt)
}

func (c *dumpContext) printBlockStmt(s *ast.BlockStmt) {
	c.nodeStart("BlockStmt")
	c.emitProp("List", s.List)
	c.nodeEnd()
}

func (c *dumpContext) printIfStmt(s *ast.IfStmt) {
	c.nodeStart("IfStmt")
	c.emitProp("Init", s.Init)
	c.emitProp("Cond", s.Cond)
	c.emitProp("Body", s.Body)
	c.emitProp("Else", s.Else)
	c.nodeEnd()
}

func (c *dumpContext) printCaseClause(stmt *ast.CaseClause) {
	// TODO:
	c.Emit("(CaseClause) %#v", stmt)
}

func (c *dumpContext) printSwitchStmt(stmt *ast.SwitchStmt) {
	// TODO:
	c.Emit("(SwitchStmt) %#v", stmt)
}

func (c *dumpContext) printTypeSwitchStmt(stmt *ast.TypeSwitchStmt) {
	// TODO:
	c.Emit("(TypeSwitchStmt) %#v", stmt)
}

func (c *dumpContext) printCommClause(stmt *ast.CommClause) {
	// TODO:
	c.Emit("(CommClause) %#v", stmt)
}

func (c *dumpContext) printSelectStmt(stmt *ast.SelectStmt) {
	// TODO:
	c.Emit("(SelectStmt) %#v", stmt)
}

func (c *dumpContext) printForStmt(stmt *ast.ForStmt) {
	// TODO:
	c.Emit("(ForStmt) %#v", stmt)
}

func (c *dumpContext) printRangeStmt(stmt *ast.RangeStmt) {
	// TODO:
	c.Emit("(RangeStmt) %#v", stmt)
}

////////////////////////////////////////////////////////////////////////////
// Decl

func (c *dumpContext) printDecl(decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.BadDecl:
		c.printBadDecl(d)
	case *ast.GenDecl:
		c.printGenDecl(d)
	case *ast.FuncDecl:
		c.printFuncDecl(d)
	default:
		c.Emit("(Decl) %#v", decl)
	}
}

func (c *dumpContext) printBadDecl(decl *ast.BadDecl) {
	// TODO:
	c.Emit("(BadDecl) %#v", decl)
}

func (c *dumpContext) printGenDecl(decl *ast.GenDecl) {
	// TODO:
	c.Emit("(GenDecl) %#v", decl)
}

func (c *dumpContext) printFuncDecl(decl *ast.FuncDecl) {
	// TODO:
	c.Emit("(FuncDecl) %#v", decl)
}

////////////////////////////////////////////////////////////////////////////
// Spec

func (c *dumpContext) printSpec(spec ast.Spec) {
	switch s := spec.(type) {
	case *ast.ImportSpec:
		c.printImportSpec(s)
	case *ast.ValueSpec:
		c.printValueSpec(s)
	case *ast.TypeSpec:
		c.printTypeSpec(s)
	default:
		c.Emit("(Spec) %#v", spec)
	}
}

func (c *dumpContext) printImportSpec(spec *ast.ImportSpec) {
	// TODO:
	c.Emit("(ImportSpec) %#v", spec)
}

func (c *dumpContext) printValueSpec(spec *ast.ValueSpec) {
	// TODO:
	c.Emit("(ValueSpec) %#v", spec)
}

func (c *dumpContext) printTypeSpec(spec *ast.TypeSpec) {
	// TODO:
	c.Emit("(TypeSpec) %#v", spec)
}

////////////////////////////////////////////////////////////////////////////
// Node

func (c *dumpContext) printNode(node ast.Node) {
	switch n := node.(type) {
	case ast.Expr:
		c.printExpr(n)
	case ast.Stmt:
		c.printStmt(n)
	case ast.Decl:
		c.printDecl(n)
	case ast.Spec:
		c.printSpec(n)
	case *ast.Comment:
		c.printComment(n)
	case *ast.CommentGroup:
		c.printCommentGroup(n)
	case *ast.Field:
		c.printField(n)
	case *ast.FieldList:
		c.printFieldList(n)
	case *ast.File:
		c.printFile(n)
	case *ast.Package:
		c.printPackage(n)
	default:
		c.Emit("(Node) %#v", node)
	}
}

func (c *dumpContext) printComment(node *ast.Comment) {
	// TODO:
	c.Emit("(Comment) %#v", node)
}

func (c *dumpContext) printCommentGroup(node *ast.CommentGroup) {
	// TODO:
	c.Emit("(CommentGroup) %#v", node)
}

func (c *dumpContext) printField(node *ast.Field) {
	// TODO:
	c.Emit("(Field) %#v", node)
}

func (c *dumpContext) printFieldList(node *ast.FieldList) {
	// TODO:
	c.Emit("(FieldList) %#v", node)
}

func (c *dumpContext) printFile(node *ast.File) {
	// TODO:
	c.Emit("(File) %#v", node)
}

func (c *dumpContext) printPackage(node *ast.Package) {
	// TODO:
	c.Emit("(Package) %#v", node)
}
