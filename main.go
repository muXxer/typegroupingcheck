package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

// LinterName is the name of the linter.
const LinterName = "typegroupingcheck"

// Analyzer is the analysis.Analyzer for the linter.
var Analyzer = &analysis.Analyzer{
	Name: LinterName,
	Doc:  "checks for omitted or grouped function parameter types",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch funcType := n.(type) {
		case *ast.FuncDecl:
			checkFunc(pass, funcType.Name.Name, funcType.Type)
		case *ast.FuncLit:
			checkFunc(pass, "inline function", funcType.Type)
		}
	})

	return nil, nil
}

func checkFunc(pass *analysis.Pass, funcName string, funcType *ast.FuncType) {
	checkFieldList := func(fl *ast.FieldList, kind string) {
		if fl == nil {
			return
		}

		for _, field := range fl.List {
			if len(field.Names) > 1 {
				pass.Reportf(field.Pos(), "grouped %s in function %q", kind, funcName)
			}
		}
	}

	checkFieldList(funcType.Params, "parameter")
	checkFieldList(funcType.Results, "result")
}

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
