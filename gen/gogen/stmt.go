package gogen

// Statement is an interface used to describe a statement(just the type of statement) in a code block
type Statement interface {
	Type() string

	AsIfStatement() *IfStatement
}

func (stmt *IfStatement) AsIfStatement() *IfStatement {
	return stmt
}
