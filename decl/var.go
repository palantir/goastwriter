// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/spec"
)

type Var struct {
	Name   string
	Type   expression.Type
	Value  astgen.ASTExpr
	Values []*spec.Value
}

func NewVar(name string, typ expression.Type) *Var {
	return &Var{
		Name: name,
		Type: typ,
	}
}

func NewVarWithValues(values ...*spec.Value) *Var {
	return &Var{
		Values: values,
	}
}

func (v *Var) ASTDecl() ast.Decl {
	var specs []ast.Spec

	if v.Values != nil {
		for _, val := range v.Values {
			specs = append(specs, val.ASTSpec())
		}
	} else {
		valueSpec := &ast.ValueSpec{
			Names: []*ast.Ident{ast.NewIdent(v.Name)},
		}
		if v.Type != "" {
			valueSpec.Type = v.Type.ToIdent()
		}
		if v.Value != nil {
			valueSpec.Values = []ast.Expr{v.Value.ASTExpr()}
		}
		specs = []ast.Spec{valueSpec}
	}

	varDecl := &ast.GenDecl{
		Tok:   token.VAR,
		Specs: specs,
	}

	if len(specs) > 1 {
		// set Lparen to non-0 value to ensure that parenthesis are rendered
		varDecl.Lparen = token.Pos(1)
	}

	return varDecl
}
