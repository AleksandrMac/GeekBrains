package task

import (
	"fmt"
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
	fmt.Println(astFile)
	return
}
