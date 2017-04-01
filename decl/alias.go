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

type Alias struct {
	Name    string
	Comment string
	Type    expression.Type
}

func (a *Alias) ASTDecl() ast.Decl {
	return &ast.GenDecl{
		Doc: comment.ToComment(a.Comment),
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(a.Name),
				Type: a.Type.ToIdent(),
			},
		},
	}
}
