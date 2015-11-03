package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", os.Stdin, 0)
	if err != nil {
		fmt.Println(err)
	}
	ast.Print(fset, f)
}
