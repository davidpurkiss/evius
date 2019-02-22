package gogen

import "go/ast"

// IfStatement defines an if statement
type IfStatement struct {
	_stmt *ast.IfStmt
}

// Type returns the type of the statement
func (stmt *IfStatement) Type() StatementType {
	return stmt.Type()
}
