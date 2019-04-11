// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/comment"
	"github.com/palantir/goastwriter/expression"
)

type Struct struct {
	Name       string
	StructType expression.StructType
	Comment    string
}

func NewStruct(name string, fields expression.StructFields, comment string) *Struct {
	return &Struct{
		Name: name,
		StructType: expression.StructType{
			Fields: fields,
		},
		Comment: comment,
	}
}

func (s *Struct) ASTDecl() ast.Decl {
	return &ast.GenDecl{
		Doc: comment.ToComment(s.Comment),
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(s.Name),
				Type: s.StructType.ASTExpr(),
			},
		},
	}
}
