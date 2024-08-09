package rule

import (
	"fmt"
	"github.com/mgechev/revive/lint"
	"go/ast"
)

// FunctionSignatureStyleRule lints given else constructs.
type FunctionSignatureStyleRule struct{}

// Apply applies the rule to given file.
func (r *FunctionSignatureStyleRule) Apply(file *lint.File, arguments lint.Arguments) []lint.Failure {
	//if len(arguments) != 1 {
	//	panic(`invalid configuration for "argument-limit"`)
	//}
	//
	//total, ok := arguments[0].(int64) // Alt. non panicking version
	//if !ok {
	//	panic(`invalid value passed as argument number to the "argument-list" rule`)
	//}

	var failures []lint.Failure

	walker := lintSignature{
		file: file,
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
	}

	ast.Walk(walker, file.AST)

	return failures
}

// Name returns the rule name.
func (r *FunctionSignatureStyleRule) Name() string {
	return "function-signature-style"
}

type lintSignature struct {
	file      *lint.File
	onFailure func(lint.Failure)
}

func (w lintSignature) Visit(n ast.Node) ast.Visitor {
	node, ok := n.(*ast.FuncDecl)
	if !ok {
		return w
	}

	for _, arg := range node.Type.Params.List {
		argStart := w.file.ToPosition(arg.Pos())
		argEnd := w.file.ToPosition(arg.End())
		fmt.Printf("arg `%s` line: %d col start: %d col end: %d\n", arg.Names[0].Name, argStart.Line, argStart.Column, argEnd.Column)
	}

	rStart := w.file.ToPosition(node.Type.Results.Pos())
	rEnd := w.file.ToPosition(node.Type.Results.End())
	fmt.Printf("Results line: %d col start: %d col end: %d\n", rStart.Line, rStart.Column, rEnd.Column)

	return w
}
