package gogen

import (
	"go/ast"
	"go/token"
)

type BinaryExpression struct {
	X     Expression
	OP    token.Token
	Y     Expression
	_expr *ast.BinaryExpr
}
