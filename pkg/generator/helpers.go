package generator

import (
	"fmt"
	"go/ast"
)

// Helper function to extract fields and their types from a struct
func extractFields(structType *ast.StructType) []FieldData {
	var fields []FieldData
	for _, field := range structType.Fields.List {
		fieldType := getFieldType(field.Type)
		for _, fieldName := range field.Names {
			fields = append(fields, FieldData{
				Name: fieldName.Name,
				Type: fieldType,
			})
		}
	}
	return fields
}

// Helper function to get field type as a string
func getFieldType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getFieldType(t.X), t.Sel.Name)
	default:
		return "unknown"
	}
}
