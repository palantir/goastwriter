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

type Const struct {
	Values []*spec.Value
}

func NewConst(values ...*spec.Value) *Const {
	return &Const{
		Values: values,
	}
}

func NewConstSingleValue(name string, typ expression.Type, value astgen.ASTExpr) *Const {
	return &Const{
		Values: []*spec.Value{
			spec.NewValue(name, typ, value),
		},
	}
}

func (c *Const) ASTDecl() ast.Decl {
	var specs []ast.Spec
	for _, val := range c.Values {
		specs = append(specs, val.ASTSpec())
	}
	constDecl := &ast.GenDecl{
		Tok:   token.CONST,
		Specs: specs,
	}
	if len(specs) > 1 {
		// set Lparen to non-0 value to ensure that parenthesis are rendered
		constDecl.Lparen = token.Pos(1)
	}
	return constDecl
}
