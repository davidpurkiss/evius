package gogen

// Expression is an interface used to describe the various types of expressions, such as binary and call expressions.
type Expression interface {
	Type() ExpressionType
}

// ExpressionType is the "enum" for various expression types and is used to identity expressions.
type ExpressionType int

const (
	// BadExpressionType can be used to identify the type of a bad expression
	BadExpressionType ExpressionType = iota
	// BinaryExpressionType can be used to identify the type of a binary expression
	BinaryExpressionType ExpressionType = iota
	// CallExpressionType can be used to identify the type of a binary expression
	CallExpressionType ExpressionType = iota
	// IndexExpressionType can be used to identify the type of an index expression
	IndexExpressionType ExpressionType = iota
	// KeyValueExpressionType can be used to identify the type of a key value expression
	KeyValueExpressionType ExpressionType = iota
	// ParenthesisExpressionType can be used to identify the type of a parenthesis expression
	ParenthesisExpressionType ExpressionType = iota
	// SelectorExpressionType can be used to identify the type of a selector expression
	SelectorExpressionType ExpressionType = iota
	// SliceExpressionType can be used to identify the type of a slice expression
	SliceExpressionType ExpressionType = iota
	// StarExpressionType can be used to identify the type of a star expression
	StarExpressionType ExpressionType = iota
	// TypeAssertExpressionType can be used to identify the type of a type assertion expression
	TypeAssertExpressionType ExpressionType = iota
	// UnaryExpressionType can be used to identify the type of a unary expression
	UnaryExpressionType ExpressionType = iota
)
