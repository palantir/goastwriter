// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
)

type FuncType struct {
	Params      FuncParams
	ReturnTypes Types
}

func (f *FuncType) ASTExpr() ast.Expr {
	return &ast.FuncType{
		Params:  f.Params.ToFieldList(),
		Results: f.ReturnTypes.ToFieldList(),
	}
}

type FuncParams []*FuncParam

func (f FuncParams) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range f {
		fields = append(fields, curr.ToASTField())
	}
	return &ast.FieldList{
		List: fields,
	}
}

type FuncParam struct {
	Names []string
	Type  Type
}

func NewFuncParam(name string, t Type) *FuncParam {
	return &FuncParam{
		Names: []string{name},
		Type:  t,
	}
}

func (f *FuncParam) ToASTField() *ast.Field {
	var names []*ast.Ident
	for _, name := range f.Names {
		names = append(names, ast.NewIdent(name))
	}
	return &ast.Field{
		Names: names,
		Type:  ast.NewIdent(string(f.Type)),
	}
}
