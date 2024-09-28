package generator

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	tmpl "github.com/fussaa/fussaa-cli/pkg/templates"
)

// Struct to store field data
type FieldData struct {
	Name string
	Type string
}

// Struct to hold template data
type MethodData struct {
	PackageName string
	StructName  string
	Fields      []FieldData
}

func GenerateMethods(packageName string, structName string, structType *ast.StructType, dest string) {
	fields := extractFields(structType)

	data := MethodData{
		PackageName: packageName,
		StructName:  structName,
		Fields:      fields,
	}

	// Generate the Go file
	outputFile, err := os.Create(filepath.Join(dest, fmt.Sprintf("%s.gmk.go", strings.ToLower(structName))))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	tmpl, err := template.New("method").Parse(tmpl.MethodTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Write the generated method to the file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

	fmt.Printf("Generated method for struct %s\n", structName)
}
