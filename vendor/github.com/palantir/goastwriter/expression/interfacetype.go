// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/comment"
)

type InterfaceType struct {
	Functions InterfaceFunctionDecls
}

func NewInterfaceType(functions ...*InterfaceFunctionDecl) *InterfaceType {
	return &InterfaceType{
		Functions: functions,
	}
}

func (i *InterfaceType) ASTExpr() ast.Expr {
	return &ast.InterfaceType{
		Methods: i.Functions.ToFieldList(),
	}
}

type InterfaceFunctionDecls []*InterfaceFunctionDecl

func (i InterfaceFunctionDecls) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range i {
		fields = append(fields, curr.ToASTField())
	}
	return &ast.FieldList{
		List: fields,
	}
}

type InterfaceFunctionDecl struct {
	Name        string
	Params      FuncParams
	ReturnTypes Types
	Comment     string
}

func (i *InterfaceFunctionDecl) ToASTField() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(i.Name),
		},
		Type: &ast.FuncType{
			Params:  i.Params.ToFieldList(),
			Results: i.ReturnTypes.ToFieldList(),
		},
		Doc: comment.ToComment(i.Comment),
	}
}
