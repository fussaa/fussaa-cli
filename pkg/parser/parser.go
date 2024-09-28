package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/fussaa/fussaa-cli/pkg/generator"
)

func ProcessFile(filenamePath string) {
	fset := token.NewFileSet()

	// Parse the Go source file
	file, err := parser.ParseFile(fset, filenamePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	packageName := file.Name.Name
	dest := filepath.Dir(filenamePath)

	// Traverse the AST and find structs
	ast.Inspect(file, func(n ast.Node) bool {
		if ts, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := ts.Type.(*ast.StructType); ok {
				structName := ts.Name.Name
				fmt.Printf("Found struct: %s\n", structName)
				generator.GenerateMethods(packageName, structName, structType, dest)
			}
		}
		return true
	})
}
