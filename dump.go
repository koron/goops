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
	for _, x := range list {
		c.printStmt(x)
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
	if list, ok := v.([]ast.Expr); ok {
		if len(list) > 0 {
			c.propStart(n)
			for _, x := range list {
				c.printExpr(x)
			}
			c.propEnd()
		}
		return
	}
	if list, ok := v.([]ast.Stmt); ok {
		if len(list) > 0 {
			c.propStart(n)
			for _, x := range list {
				c.printStmt(x)
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
	switch x := expr.(type) {
	case *ast.BadExpr:
		c.printBadExpr(x)
	case *ast.Ident:
		c.printIdent(x)
	case *ast.Ellipsis:
		c.printEllipsis(x)
	case *ast.BasicLit:
		c.printBasicLit(x)
	case *ast.FuncLit:
		c.printFuncLit(x)
	case *ast.CompositeLit:
		c.printCompositeLit(x)
	case *ast.ParenExpr:
		c.printParenExpr(x)
	case *ast.SelectorExpr:
		c.printSelectorExpr(x)
	case *ast.IndexExpr:
		c.printIndexExpr(x)
	case *ast.SliceExpr:
		c.printSliceExpr(x)
	case *ast.TypeAssertExpr:
		c.printTypeAssertExpr(x)
	case *ast.CallExpr:
		c.printCallExpr(x)
	case *ast.StarExpr:
		c.printStarExpr(x)
	case *ast.UnaryExpr:
		c.printUnaryExpr(x)
	case *ast.BinaryExpr:
		c.printBinaryExpr(x)
	case *ast.KeyValueExpr:
		c.printKeyValueExpr(x)
	case *ast.ArrayType:
		c.printArrayType(x)
	case *ast.StructType:
		c.printStructType(x)
	case *ast.FuncType:
		c.printFuncType(x)
	case *ast.InterfaceType:
		c.printInterfaceType(x)
	case *ast.MapType:
		c.printMapType(x)
	case *ast.ChanType:
		c.printChanType(x)
	default:
		c.Emit("(Expr) %#v", expr)
	}
}

func (c *dumpContext) printBadExpr(x *ast.BadExpr) {
	c.nodeStart("BadExpr")
	c.nodeEnd()
}

func (c *dumpContext) printIdent(x *ast.Ident) {
	c.nodeStart("Ident")
	c.emitProp("Name", x.Name)
	c.emitProp("Obj", x.Obj)
	c.nodeEnd()
}

func (c *dumpContext) printEllipsis(x *ast.Ellipsis) {
	c.nodeStart("Ellipsis")
	c.emitProp("Elt", x.Elt)
	c.nodeEnd()
}

func (c *dumpContext) printBasicLit(x *ast.BasicLit) {
	c.nodeStart("BasicLit")
	c.emitProp("Kind", x.Kind)
	c.emitProp("Value", x.Value)
	c.nodeEnd()
}

func (c *dumpContext) printFuncLit(x *ast.FuncLit) {
	c.nodeStart("FuncLit")
	c.emitProp("Type", x.Type)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printCompositeLit(x *ast.CompositeLit) {
	c.nodeStart("CompositeLit")
	c.emitProp("Type", x.Type)
	c.emitProp("Elts", x.Elts)
	c.nodeEnd()
}

func (c *dumpContext) printParenExpr(x *ast.ParenExpr) {
	c.nodeStart("ParenExpr")
	c.emitProp("X", x.X)
	c.nodeEnd()
}

func (c *dumpContext) printSelectorExpr(x *ast.SelectorExpr) {
	c.nodeStart("SelectorExpr")
	c.emitProp("X", x.X)
	c.emitProp("Sel", x.Sel)
	c.nodeEnd()
}

func (c *dumpContext) printIndexExpr(x *ast.IndexExpr) {
	c.nodeStart("IndexExpr")
	c.emitProp("X", x.X)
	c.emitProp("Index", x.Index)
	c.nodeEnd()
}

func (c *dumpContext) printSliceExpr(x *ast.SliceExpr) {
	c.nodeStart("SliceExpr")
	c.emitProp("X", x.X)
	c.emitProp("Low", x.Low)
	c.emitProp("High", x.High)
	c.emitProp("Max", x.Max)
	c.emitProp("Slice3", x.Slice3)
	c.nodeEnd()
}

func (c *dumpContext) printTypeAssertExpr(x *ast.TypeAssertExpr) {
	c.nodeStart("TypeAssertExpr")
	c.emitProp("X", x.X)
	c.emitProp("Type", x.Type)
	c.nodeEnd()
}

func (c *dumpContext) printCallExpr(x *ast.CallExpr) {
	c.nodeStart("CallExpr")
	c.emitProp("Func", x.Fun)
	c.emitProp("Args", x.Args)
	c.nodeEnd()
}

func (c *dumpContext) printStarExpr(x *ast.StarExpr) {
	c.nodeStart("StarExpr")
	c.emitProp("X", x.X)
	c.nodeEnd()
}

func (c *dumpContext) printUnaryExpr(x *ast.UnaryExpr) {
	c.nodeStart("UnaryExpr")
	c.emitProp("Op", x.Op)
	c.emitProp("X", x.X)
	c.nodeEnd()
}

func (c *dumpContext) printBinaryExpr(x *ast.BinaryExpr) {
	c.nodeStart("BinaryExpr")
	c.emitProp("X", x.X)
	c.emitProp("Op", x.Op)
	c.emitProp("Y", x.Y)
	c.nodeEnd()
}

func (c *dumpContext) printKeyValueExpr(x *ast.KeyValueExpr) {
	c.nodeStart("KeyValueExpr")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	c.nodeEnd()
}

func (c *dumpContext) printArrayType(x *ast.ArrayType) {
	c.nodeStart("ArrayType")
	c.emitProp("Len", x.Len)
	c.emitProp("Elt", x.Elt)
	c.nodeEnd()
}

func (c *dumpContext) printStructType(x *ast.StructType) {
	c.nodeStart("StructType")
	c.emitProp("FieldList", x.Fields)
	c.emitProp("Incomplete", x.Incomplete)
	c.nodeEnd()
}

func (c *dumpContext) printFuncType(x *ast.FuncType) {
	c.nodeStart("FuncType")
	c.emitProp("Params", x.Params)
	c.emitProp("Results", x.Results)
	c.nodeEnd()
}

func (c *dumpContext) printInterfaceType(x *ast.InterfaceType) {
	c.nodeStart("InterfaceType")
	c.emitProp("Methods", x.Methods)
	c.emitProp("Incomplete", x.Incomplete)
	c.nodeEnd()
}

func (c *dumpContext) printMapType(x *ast.MapType) {
	c.nodeStart("MapType")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	c.nodeEnd()
}

func (c *dumpContext) printChanType(x *ast.ChanType) {
	c.nodeStart("ChanType")
	c.emitProp("Dir", x.Dir)
	c.emitProp("Value", x.Value)
	c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Stmt

func (c *dumpContext) printStmt(stmt ast.Stmt) {
	switch x := stmt.(type) {
	case *ast.BadStmt:
		c.printBadStmt(x)
	case *ast.DeclStmt:
		c.printDeclStmt(x)
	case *ast.EmptyStmt:
		c.printEmptyStmt(x)
	case *ast.LabeledStmt:
		c.printLabeledStmt(x)
	case *ast.ExprStmt:
		c.printExprStmt(x)
	case *ast.SendStmt:
		c.printSendStmt(x)
	case *ast.IncDecStmt:
		c.printIncDecStmt(x)
	case *ast.AssignStmt:
		c.printAssignStmt(x)
	case *ast.GoStmt:
		c.printGoStmt(x)
	case *ast.DeferStmt:
		c.printDeferStmt(x)
	case *ast.ReturnStmt:
		c.printReturnStmt(x)
	case *ast.BranchStmt:
		c.printBranchStmt(x)
	case *ast.BlockStmt:
		c.printBlockStmt(x)
	case *ast.IfStmt:
		c.printIfStmt(x)
	case *ast.CaseClause:
		c.printCaseClause(x)
	case *ast.SwitchStmt:
		c.printSwitchStmt(x)
	case *ast.TypeSwitchStmt:
		c.printTypeSwitchStmt(x)
	case *ast.CommClause:
		c.printCommClause(x)
	case *ast.SelectStmt:
		c.printSelectStmt(x)
	case *ast.ForStmt:
		c.printForStmt(x)
	case *ast.RangeStmt:
		c.printRangeStmt(x)
	default:
		c.Emit("(Stmt) %#v", stmt)
	}
}

func (c *dumpContext) printBadStmt(x *ast.BadStmt) {
	c.nodeStart("BadStmt")
	c.nodeEnd()
}

func (c *dumpContext) printDeclStmt(x *ast.DeclStmt) {
	c.nodeStart("DeclStmt")
	c.emitProp("Decl", x.Decl)
	c.nodeEnd()
}

func (c *dumpContext) printEmptyStmt(x *ast.EmptyStmt) {
	c.nodeStart("EmptyStmt")
	c.nodeEnd()
}

func (c *dumpContext) printLabeledStmt(x *ast.LabeledStmt) {
	c.nodeStart("LabeledStmt")
	c.emitProp("Label", x.Label)
	c.emitProp("Stmt", x.Stmt)
	c.nodeEnd()
}

func (c *dumpContext) printExprStmt(x *ast.ExprStmt) {
	c.nodeStart("ExprStmt")
	c.emitProp("X", x.X)
	c.nodeEnd()
}

func (c *dumpContext) printSendStmt(x *ast.SendStmt) {
	c.nodeStart("SendStmt")
	c.emitProp("Chan", x.Chan)
	c.emitProp("Value", x.Value)
	c.nodeEnd()
}

func (c *dumpContext) printIncDecStmt(x *ast.IncDecStmt) {
	c.nodeStart("IncDecStmt")
	c.emitProp("X", x.X)
	c.emitProp("Tok", x.Tok)
	c.nodeEnd()
}

func (c *dumpContext) printAssignStmt(x *ast.AssignStmt) {
	c.nodeStart("AssignStmt")
	c.emitProp("Lhs", x.Lhs)
	c.emitProp("Tok", x.Tok)
	c.emitProp("Rhs", x.Rhs)
	c.nodeEnd()
}

func (c *dumpContext) printGoStmt(x *ast.GoStmt) {
	c.nodeStart("GoStmt")
	c.emitProp("Call", x.Call)
	c.nodeEnd()
}

func (c *dumpContext) printDeferStmt(x *ast.DeferStmt) {
	c.nodeStart("DeferStmt")
	c.emitProp("Call", x.Call)
	c.nodeEnd()
}

func (c *dumpContext) printReturnStmt(x *ast.ReturnStmt) {
	c.nodeStart("ReturnStmt")
	c.emitProp("Results", x.Results)
	c.nodeEnd()
}

func (c *dumpContext) printBranchStmt(x *ast.BranchStmt) {
	c.nodeStart("BranchStmt")
	c.emitProp("Tok", x.Tok)
	c.emitProp("Label", x.Label)
	c.nodeEnd()
}

func (c *dumpContext) printBlockStmt(x *ast.BlockStmt) {
	c.nodeStart("BlockStmt")
	c.emitProp("List", x.List)
	c.nodeEnd()
}

func (c *dumpContext) printIfStmt(x *ast.IfStmt) {
	c.nodeStart("IfStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Cond", x.Cond)
	c.emitProp("Body", x.Body)
	c.emitProp("Else", x.Else)
	c.nodeEnd()
}

func (c *dumpContext) printCaseClause(x *ast.CaseClause) {
	c.nodeStart("CaseClause")
	c.emitProp("List", x.List)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printSwitchStmt(x *ast.SwitchStmt) {
	c.nodeStart("SwitchStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Tag", x.Tag)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printTypeSwitchStmt(x *ast.TypeSwitchStmt) {
	c.nodeStart("TypeSwitchStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Assign", x.Assign)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printCommClause(x *ast.CommClause) {
	c.nodeStart("CommClause")
	c.emitProp("Comm", x.Comm)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printSelectStmt(x *ast.SelectStmt) {
	c.nodeStart("SelectStmt")
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printForStmt(x *ast.ForStmt) {
	c.nodeStart("ForStmt")
	c.emitProp("Init", x.Init)
	c.emitProp("Cond", x.Cond)
	c.emitProp("Post", x.Post)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

func (c *dumpContext) printRangeStmt(x *ast.RangeStmt) {
	c.nodeStart("RangeStmt")
	c.emitProp("Key", x.Key)
	c.emitProp("Value", x.Value)
	c.emitProp("Tok", x.Tok)
	c.emitProp("X", x.X)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Decl

func (c *dumpContext) printDecl(decl ast.Decl) {
	switch x := decl.(type) {
	case *ast.BadDecl:
		c.printBadDecl(x)
	case *ast.GenDecl:
		c.printGenDecl(x)
	case *ast.FuncDecl:
		c.printFuncDecl(x)
	default:
		c.Emit("(Decl) %#v", decl)
	}
}

func (c *dumpContext) printBadDecl(x *ast.BadDecl) {
	c.nodeStart("BadDecl")
	c.nodeEnd()
}

func (c *dumpContext) printGenDecl(x *ast.GenDecl) {
	c.nodeStart("GenDecl")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Tok", x.Tok)
	c.emitProp("Specs", x.Specs)
	c.nodeEnd()
}

func (c *dumpContext) printFuncDecl(x *ast.FuncDecl) {
	c.nodeStart("FuncDecl")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Recv", x.Recv)
	c.emitProp("Name", x.Name)
	c.emitProp("Type", x.Type)
	c.emitProp("Body", x.Body)
	c.nodeEnd()
}

////////////////////////////////////////////////////////////////////////////
// Spec

func (c *dumpContext) printSpec(spec ast.Spec) {
	switch x := spec.(type) {
	case *ast.ImportSpec:
		c.printImportSpec(x)
	case *ast.ValueSpec:
		c.printValueSpec(x)
	case *ast.TypeSpec:
		c.printTypeSpec(x)
	default:
		c.Emit("(Spec) %#v", spec)
	}
}

func (c *dumpContext) printImportSpec(x *ast.ImportSpec) {
	c.nodeStart("ImportSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Path", x.Path)
	c.emitProp("Comment", x.Comment)
	c.nodeEnd()
}

func (c *dumpContext) printValueSpec(x *ast.ValueSpec) {
	c.nodeStart("ValueSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Names", x.Names)
	c.emitProp("Type", x.Type)
	c.emitProp("Values", x.Values)
	c.emitProp("Comment", x.Comment)
}

func (c *dumpContext) printTypeSpec(x *ast.TypeSpec) {
	c.nodeStart("TypeSpec")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Type", x.Type)
	c.emitProp("Comment", x.Comment)
}

////////////////////////////////////////////////////////////////////////////
// Node

func (c *dumpContext) printNode(node ast.Node) {
	switch x := node.(type) {
	case ast.Expr:
		c.printExpr(x)
	case ast.Stmt:
		c.printStmt(x)
	case ast.Decl:
		c.printDecl(x)
	case ast.Spec:
		c.printSpec(x)
	case *ast.Comment:
		c.printComment(x)
	case *ast.CommentGroup:
		c.printCommentGroup(x)
	case *ast.Field:
		c.printField(x)
	case *ast.FieldList:
		c.printFieldList(x)
	case *ast.File:
		c.printFile(x)
	case *ast.Package:
		c.printPackage(x)
	default:
		c.Emit("(Node) %#v", node)
	}
}

func (c *dumpContext) printComment(x *ast.Comment) {
	c.nodeStart("Comment")
	c.emitProp("Text", x.Text)
	c.nodeEnd()
}

func (c *dumpContext) printCommentGroup(x *ast.CommentGroup) {
	c.nodeStart("CommentGroup")
	c.emitProp("List", x.List)
	c.nodeEnd()
}

func (c *dumpContext) printField(x *ast.Field) {
	c.nodeStart("Field")
	c.emitProp("Doc", x.Doc)
	c.emitProp("Names", x.Names)
	c.emitProp("Type", x.Type)
	c.emitProp("Tag", x.Tag)
	c.emitProp("Comment", x.Comment)
	c.nodeEnd()
}

func (c *dumpContext) printFieldList(x *ast.FieldList) {
	c.nodeStart("FieldList")
	c.emitProp("List", x.List)
	c.nodeEnd()
}

func (c *dumpContext) printFile(x *ast.File) {
	c.nodeStart("File")
	c.emitProp("Doc",x.Doc)
	c.emitProp("Name", x.Name)
	c.emitProp("Decls", x.Decls)
	c.emitProp("Scope", x.Scope)
	c.emitProp("Imports", x.Imports)
	c.emitProp("Unresolved", x.Unresolved)
	c.emitProp("Comments", x.Comments)
	c.nodeEnd()
}

func (c *dumpContext) printPackage(x *ast.Package) {
	c.nodeStart("Package")
	c.emitProp("Name", x.Name)
	c.emitProp("Scope", x.Scope)
	c.emitProp("Imports", x.Imports)
	c.emitProp("Files", x.Files)
	c.nodeEnd()
}
