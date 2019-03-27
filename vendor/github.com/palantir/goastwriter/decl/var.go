// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/expression"
)

type Var struct {
	Name string
	Type expression.Type
}

func NewVar(name string, typ expression.Type) *Var {
	return &Var{
		Name: name,
		Type: typ,
	}
}

func (v *Var) ASTDecl() ast.Decl {
	return &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					ast.NewIdent(v.Name),
				},
				Type: &ast.Ident{
					Name: string(v.Type),
				},
			},
		},
	}
}
