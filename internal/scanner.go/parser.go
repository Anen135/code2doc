package scanner

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Summary struct {
	FileName string
	Methods  []string
}

func AnalyzeCode(fileName string, content string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, content, parser.ParseComments)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Файл: %s\n\n", fileName))

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			sb.WriteString(fmt.Sprintf("Функция: %s\n", x.Name.Name))
			if x.Doc != nil {
				sb.WriteString(fmt.Sprintf("Описание: %s", x.Doc.Text()))
			}
		case *ast.TypeSpec:
			sb.WriteString(fmt.Sprintf("Тип: %s\n", x.Name.Name))
		}
		return true
	})

	return sb.String(), nil
}
