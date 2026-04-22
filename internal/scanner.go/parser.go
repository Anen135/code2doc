package scanner

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"code2doc/internal/logger"
)

type Summary struct {
	FileName string
	Methods  []string
}

func AnalyzeCode(fileName string, content string) (string, error) {
	logger.Debug(fmt.Sprintf("Parsing file: %s", fileName))

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, content, parser.ParseComments)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to parse file %s: %v", fileName, err))
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Файл: %s\n\n", fileName))

	funcCount := 0
	typeCount := 0

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			sb.WriteString(fmt.Sprintf("Функция: %s\n", x.Name.Name))
			if x.Doc != nil {
				sb.WriteString(fmt.Sprintf("Описание: %s", x.Doc.Text()))
			}
			funcCount++
		case *ast.TypeSpec:
			sb.WriteString(fmt.Sprintf("Тип: %s\n", x.Name.Name))
			typeCount++
		}
		return true
	})

	logger.Info(fmt.Sprintf("Parsed %s: found %d functions, %d types", fileName, funcCount, typeCount))

	return sb.String(), nil
}
