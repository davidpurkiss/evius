package gogen

// Statement is an interface used to describe a statement(just the type of statement) in a code block
type Statement interface {
	Type() StatementType

	AsIfStatement() *IfStatement
	AsForStatement() *ForStatement
}

// StatementType is the "enum" for various statement types and is used to identity statements.
type StatementType int

const (
	// AssignStatementType identifies an assignment statement
	AssignStatementType StatementType = iota
	// BadStatementType identifies a bad statement
	BadStatementType StatementType = iota
	// BlockStatementType identifies a block statement
	BlockStatementType StatementType = iota
	// BranchStatementType identifies a branch statement
	BranchStatementType StatementType = iota
	// DeferStatementType identifies a defer statement
	DeferStatementType StatementType = iota
	// EmptyStatementType identifies an empty statement
	EmptyStatementType StatementType = iota
	// ExpressionStatementType identifies an expression statement
	ExpressionStatementType StatementType = iota
	// ForStatementType identifies a for statement
	ForStatementType StatementType = iota
	// GoStatementType identifies a go statement
	GoStatementType StatementType = iota
	// IfStatementType identifies an if statement
	IfStatementType StatementType = iota
	// IncrementDecrementStatementType identifies a increment decrement statement
	IncrementDecrementStatementType StatementType = iota
	// LabeledStatementType identifies a labeled statement
	LabeledStatementType StatementType = iota
	// RangeStatementType identifies a range statement
	RangeStatementType StatementType = iota
	// ReturnStatementType identifies a return statement
	ReturnStatementType StatementType = iota
	// SelectStatementType identifies a select statement
	SelectStatementType StatementType = iota
	// SendStatementType identifies a send statement
	SendStatementType StatementType = iota
	// SwitchStatementType identifies a switch statement
	SwitchStatementType StatementType = iota
	// TypeSwitchStatementType identifies a type switch statement
	TypeSwitchStatementType StatementType = iota
)

// AsIfStatement returns an if statement from a statement if it is an IfStatement
func (stmt *IfStatement) AsIfStatement() *IfStatement {
	return stmt
}

// AsForStatement returns a for statement from a statement if it is a ForStatement
func (stmt *IfStatement) AsForStatement() *ForStatement {
	return nil
}
