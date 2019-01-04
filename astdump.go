package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("\tastdump [filename]")
		os.Exit(1)
	}
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("cannot open file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, file, 0)
	if err != nil {
		fmt.Println(err)
	}
	ast.Print(fset, f)
}
