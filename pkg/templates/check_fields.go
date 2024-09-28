package template

import "fmt"

// Template for generating methods
var MethodTemplate = fmt.Sprintf(`%s
package {{.PackageName}}

import "fmt"

// CheckFields method generated based on struct fields.
func (s *{{.StructName}}) CheckFields() {
    {{range .Fields}}
    // Field: {{.Name}} (Type: {{.Type}})
    {{if eq .Type "string"}}
    if s.{{.Name}} == "" {
        fmt.Println("Field '{{.Name}}' is an empty string.")
    } else {
        fmt.Println("Field '{{.Name}}' has value:", s.{{.Name}})
    }
    {{else if eq .Type "int"}}
    if s.{{.Name}} == 0 {
        fmt.Println("Field '{{.Name}}' is zero.")
    } else {
        fmt.Println("Field '{{.Name}}' has value:", s.{{.Name}})
    }
    {{end}}
    {{end}}
}
`, defaultHeader)
