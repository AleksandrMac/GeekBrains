package task

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Task2 - Написать функцию, которая принимает на вход имя файла и название функции.
// Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
// Результат работы должен возвращать количество вызовов int и ошибку error.
// Разрешается использовать только go/parser, go/ast и go/token.
func Task2(fileName string, funcName string) (count int32, err error) {
	fset := token.NewFileSet()
	// парсим файл, чтобы получить AST
	astFile, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return
	}
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		count += funcNameFromStmt(funcDecl.Body, funcName)

		if funcDecl.Name.String() == funcName {
			count++
		}

	}
	return
}

func funcNameFromStmt(stmt ast.Stmt, funcName string) (count int32) {
	switch stmt.(type) {
	case *ast.BlockStmt:
		for _, block := range stmt.(*ast.BlockStmt).List {
			count += funcNameFromStmt(block, funcName)
		}
	case *ast.ExprStmt:
		count += funcNameFromExpr(stmt.(*ast.ExprStmt).X, funcName)
	case *ast.IfStmt:
		count += funcNameFromExpr(stmt.(*ast.IfStmt).Cond, funcName)
		count += funcNameFromStmt(stmt.(*ast.IfStmt).Body, funcName)
		if stmt.(*ast.IfStmt).Else != nil {
			count += funcNameFromStmt(stmt.(*ast.IfStmt).Else, funcName)
		}
	case *ast.ForStmt:
		count += funcNameFromStmt(stmt.(*ast.ForStmt).Init, funcName)
		count += funcNameFromStmt(stmt.(*ast.ForStmt).Body, funcName)
		count += funcNameFromExpr(stmt.(*ast.ForStmt).Cond, funcName)
	case *ast.AssignStmt:
		for _, assign := range stmt.(*ast.AssignStmt).Rhs {
			count += funcNameFromExpr(assign, funcName)
		}
	case *ast.RangeStmt:
		count += funcNameFromExpr(stmt.(*ast.RangeStmt).X, funcName)
		count += funcNameFromStmt(stmt.(*ast.RangeStmt).Body, funcName)
	case *ast.ReturnStmt:
		for _, ret := range stmt.(*ast.ReturnStmt).Results {
			count += funcNameFromExpr(ret, funcName)
		}
	}
	return
}

func funcNameFromExpr(e ast.Expr, funcName string) (count int32) {
	switch e.(type) {
	case *ast.BinaryExpr:
		t := e.(*ast.BinaryExpr)
		count += funcNameFromExpr(t.X, funcName)
		count += funcNameFromExpr(t.Y, funcName)
	case *ast.Ident:
		t := e.(*ast.Ident)
		if t.Name == funcName {
			count++
		}
	case *ast.CallExpr:
		t := e.(*ast.CallExpr)
		count += funcNameFromExpr(t.Fun, funcName)
		for _, arg := range t.Args {
			count += funcNameFromExpr(arg, funcName)
		}
	case *ast.SelectorExpr:
		t := e.(*ast.SelectorExpr)
		count += funcNameFromExpr(t.Sel, funcName)
	case *ast.BasicLit:
		t := e.(*ast.BasicLit)
		if t.Value == funcName {
			count++
		}
	}

	return
}
