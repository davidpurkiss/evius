package gogen

import (
	"go/ast"
	"go/token"
)

// BinaryExpression represents any binary expression
type BinaryExpression struct {
	X    Expression
	OP   token.Token
	Y    Expression
	expr *ast.BinaryExpr
}

// Type returns the type of this expression (useful when you only have the interface)
func (expr BinaryExpression) Type() ExpressionType {
	return BinaryExpressionType
}
