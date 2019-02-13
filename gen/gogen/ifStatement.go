package gogen

import "go/ast"

type IfStatement struct {
	_stmt *ast.IfStmt
}

func (stmt *IfStatement) Type() string {
	return stmt.Type()
}

func (stmt *IfStatement) AsIfStatement() *IfStatement {
	return stmt
}
