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
func Task2(fileName, funcName string) (int32, error) {
	fset := token.NewFileSet()
	// парсим файл, чтобы получить AST
	astFile, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		return 0, nil
	}
	var count int32
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
	return count, nil
}

func funcNameFromStmt(stmt ast.Stmt, funcName string) (count int32) {
	switch v := stmt.(type) {
	case *ast.BlockStmt:
		for _, block := range v.List {
			count += funcNameFromStmt(block, funcName)
		}
	case *ast.ExprStmt:
		count += funcNameFromExpr(v.X, funcName)
	case *ast.IfStmt:
		count += funcNameFromExpr(v.Cond, funcName)
		count += funcNameFromStmt(v.Body, funcName)
		if stmt.(*ast.IfStmt).Else != nil {
			count += funcNameFromStmt(v.Else, funcName)
		}
	case *ast.ForStmt:
		count += funcNameFromStmt(v.Init, funcName)
		count += funcNameFromStmt(v.Body, funcName)
		count += funcNameFromExpr(v.Cond, funcName)
	case *ast.AssignStmt:
		for _, assign := range v.Rhs {
			count += funcNameFromExpr(assign, funcName)
		}
	case *ast.RangeStmt:
		count += funcNameFromExpr(v.X, funcName)
		count += funcNameFromStmt(v.Body, funcName)
	case *ast.ReturnStmt:
		for _, ret := range v.Results {
			count += funcNameFromExpr(ret, funcName)
		}
	}
	return count
}

func funcNameFromExpr(e ast.Expr, funcName string) (count int32) {
	switch v := e.(type) {
	case *ast.BinaryExpr:
		count += funcNameFromExpr(v.X, funcName)
		count += funcNameFromExpr(v.Y, funcName)
	case *ast.Ident:
		if v.Name == funcName {
			count++
		}
	case *ast.CallExpr:
		count += funcNameFromExpr(v.Fun, funcName)
		for _, arg := range v.Args {
			count += funcNameFromExpr(arg, funcName)
		}
	case *ast.SelectorExpr:
		count += funcNameFromExpr(v.Sel, funcName)
	case *ast.BasicLit:
		if v.Value == funcName {
			count++
		}
	}
	return count
}
