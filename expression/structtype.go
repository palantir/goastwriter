// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/comment"
)

type StructType struct {
	Fields StructFields
}

func NewStructType(fields ...*StructField) *StructType {
	return &StructType{
		Fields: fields,
	}
}

func (s *StructType) ASTExpr() ast.Expr {
	return &ast.StructType{
		Fields: s.Fields.ToFieldList(),
	}
}

type StructField struct {
	Name    string
	Type    Type
	Tag     string
	Comment string
}

type StructFields []*StructField

func (s StructFields) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range s {
		fields = append(fields, curr.ToASTField())
	}
	return &ast.FieldList{
		List: fields,
	}
}

func (f *StructField) ToASTField() *ast.Field {
	var tag *ast.BasicLit
	if f.Tag != "" {
		tag = &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("`%s`", f.Tag),
		}
	}

	return &ast.Field{
		Doc: comment.ToComment(f.Comment),
		Names: []*ast.Ident{
			ast.NewIdent(f.Name),
		},
		Type: f.Type.ToIdent(),
		Tag:  tag,
	}
}
