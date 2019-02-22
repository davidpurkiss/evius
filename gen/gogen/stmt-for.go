package gogen

// ForStatement defines a for loop statement
type ForStatement struct {
	Init Statement
	Cond Expression
	Post Statement
	Body Statement
}

// Type returns the type of the statement
func (stmt *ForStatement) Type() StatementType {
	return stmt.Type()
}
