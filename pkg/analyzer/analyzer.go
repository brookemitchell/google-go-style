package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "googlestyle",
	Doc:      "Checks that code follows Google's Go Style Guide naming conventions (no snake_case)",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

// isSnakeCase checks if a string follows snake_case naming convention
// or has underscores in it (which we want to flag).
func isSnakeCase(s string) bool {
	// Flag any identifier with underscores
	return strings.Contains(s, "_") && s[0] != '_'
}

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.TypeSpec)(nil),
		(*ast.ValueSpec)(nil),
		(*ast.StructType)(nil),
		(*ast.AssignStmt)(nil),
	}

	insp.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.FuncDecl:
			checkNamingConvention(pass, n.Name.Name, n.Pos())
		case *ast.TypeSpec:
			checkNamingConvention(pass, n.Name.Name, n.Pos())
		case *ast.ValueSpec:
			for _, name := range n.Names {
				checkNamingConvention(pass, name.Name, name.Pos())
			}
		case *ast.StructType:
			// Check struct field names
			for _, field := range n.Fields.List {
				for _, name := range field.Names {
					checkNamingConvention(pass, name.Name, name.Pos())
				}
			}
		case *ast.AssignStmt:
			// Check for local variable declarations (:= syntax)
			if n.Tok == token.DEFINE {
				for _, expr := range n.Lhs {
					if ident, ok := expr.(*ast.Ident); ok {
						checkNamingConvention(pass, ident.Name, ident.Pos())
					}
				}
			}
		}
	})

	return nil, nil
}

func checkNamingConvention(pass *analysis.Pass, name string, pos token.Pos) {
	// Skip one-letter names or names starting with underscores
	if len(name) <= 1 || name[0] == '_' {
		return
	}

	if isSnakeCase(name) {
		pass.Reportf(pos, "snake_case name '%s', use MixedCaps or mixedCaps rather than underscore when writing multi-word names.", name)
	}
}
