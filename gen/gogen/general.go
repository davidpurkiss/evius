package gogen

import (
	"fmt"
	"go/ast"
)

func getCommentGroup(text string) *ast.CommentGroup {
	var cg *ast.CommentGroup
	if text != "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Text: fmt.Sprint("// ", text), Slash: 0}},
		}
	}

	return cg
}
